package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/altipla-consulting/database"
	"github.com/altipla-consulting/redis"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/hcl"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(errors.ErrorStack(err))
	}
}

type Field struct {
	Name        string
	Values      []interface{}
	Placeholder string
}

func run() error {
	if err := fillSQL(); err != nil {
		return errors.Trace(err)
	}

	if err := fillRedis(); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func fillSQL() error {
	dbs, err := ioutil.ReadDir("data/sql")
	if err != nil {
		return errors.Trace(err)
	}

	for _, db := range dbs {
		tables, err := ioutil.ReadDir(fmt.Sprintf("data/sql/%s", db.Name()))
		if err != nil {
			return errors.Trace(err)
		}

		credentials := database.Credentials{
			User:      "dev-user",
			Password:  "dev-password",
			Address:   "database:3306",
			Database:  db.Name(),
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		}
		sess, err := sql.Open("mysql", credentials.String())
		if err != nil {
			return errors.Trace(err)
		}
		sess.SetMaxOpenConns(1)
		sess.SetMaxIdleConns(0)

		for _, table := range tables {
			collection := table.Name()
			collection = collection[:len(collection)-len(filepath.Ext(table.Name()))]

			// Evita el prefijo de orden que simplemente usamos por las relaciones.
			parts := strings.Split(collection, "-")
			if len(parts) > 1 {
				collection = parts[1]
			}

			f, err := os.Open(fmt.Sprintf("data/sql/%s/%s", db.Name(), table.Name()))
			if err != nil {
				return errors.Trace(err)
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				return errors.Trace(err)
			}
			strContent := string(content)

			data := map[string]interface{}{}
			if err := hcl.Decode(&data, strContent); err != nil {
				return errors.Trace(err)
			}

			// Carga la lista de entidades.
			if len(data) != 1 {
				return errors.New("there is more than one root key, one table per file only")
			}
			if _, ok := data[collection]; !ok {
				return errors.Errorf("table name not found in the data file, expected: %s", collection)
			}
			entities, ok := data[collection].([]map[string]interface{})
			if !ok {
				return errors.New("cannot find the list of entities")
			}

			log.WithFields(log.Fields{
				"db":         db.Name(),
				"collection": collection,
				"entities":   len(entities),
			}).Info("Fill data")

			// Valida otros aspectos de alto nivel del fichero de datos.
			if strings.Contains(strContent, "= {") || strings.Contains(strContent, "={") {
				return errors.New("use the maps directly without assignment")
			}

			for _, entity := range entities {
				var fields []*Field
				var pk *Field
				for k, v := range entity {
					if k == "revision" {
						return errors.Errorf("do not specify revision in the data")
					}

					field := &Field{
						Name:        k,
						Values:      []interface{}{v},
						Placeholder: "?",
					}
					fields = append(fields, field)

					if k == "code" || k == "id" {
						log.WithField("pk", v).Info("... entity")
						pk = field
					}

					if strings.HasSuffix(k, ".point") {
						vt := v.([]map[string]interface{})
						if len(vt) != 1 {
							return errors.Errorf("unexpected column type: %T %#v", v, v)
						}

						field.Name = strings.TrimSuffix(k, ".point")
						field.Placeholder = "POINT(?, ?)"
						field.Values = []interface{}{vt[0]["lng"], vt[0]["lat"]}
					} else {
						switch vt := v.(type) {
						// Valores primitivos.
						case string, int64, float64, bool, int:
							// empty

						// Traducciones y valores de proveedor.
						case []map[string]interface{}:
							if len(vt) != 1 {
								return errors.Errorf("unexpected column type: %T --- %#v", v, v)
							}

							var buf bytes.Buffer
							if err := json.NewEncoder(&buf).Encode(vt[0]); err != nil {
								return errors.Trace(err)
							}

							field.Values = []interface{}{buf.String()}

						// Listas de strings o números que se serializan como JSON.
						case []interface{}:
							var buf bytes.Buffer
							if err := json.NewEncoder(&buf).Encode(v); err != nil {
								return errors.Trace(err)
							}

							field.Values = []interface{}{buf.String()}

						default:
							return errors.Errorf("unexpected column type: %T %#v", v, v)
						}
					}
				}

				fields = append(fields, &Field{
					Name:        "revision",
					Placeholder: "?",
					Values:      []interface{}{0},
				})

				var n int64
				sql := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", collection, pk.Name)
				if err := sess.QueryRow(sql, pk.Values[0]).Scan(&n); err != nil {
					return errors.Trace(err)
				}

				if n == 0 {
					cols := []string{}
					placeholders := []string{}
					values := []interface{}{}
					for _, field := range fields {
						cols = append(cols, field.Name)
						placeholders = append(placeholders, field.Placeholder)
						values = append(values, field.Values...)
					}
					sql := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (%s)`, collection, strings.Join(cols, ","), strings.Join(placeholders, ","))
					if _, err := sess.Exec(sql, values...); err != nil {
						log.WithFields(log.Fields{"sql": sql, "values": values}).Error("INSERT failed")
						return errors.Trace(err)
					}
				} else {
					values := []interface{}{}
					cols := []string{}
					for _, field := range fields {
						cols = append(cols, fmt.Sprintf("%s = %s", field.Name, field.Placeholder))
						values = append(values, field.Values...)
					}
					values = append(values, pk.Values[0])

					sql := fmt.Sprintf(`UPDATE %s SET %s WHERE %s = ?`, collection, strings.Join(cols, ","), pk.Name)
					if _, err := sess.Exec(sql, values...); err != nil {
						log.WithFields(log.Fields{"sql": sql, "values": values}).Error("UPDATE failed")
						return errors.Trace(err)
					}
				}
			}
		}
	}

	return nil
}

func fillRedis() error {
	apps, err := ioutil.ReadDir("data/redis")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return errors.Trace(err)
	}

	for _, app := range apps {
		sess := redis.Open("redis:6379", app.Name())

		keys, err := ioutil.ReadDir(fmt.Sprintf("data/redis/%s", app.Name()))
		if err != nil {
			return errors.Trace(err)
		}

		for _, key := range keys {
			keyName := key.Name()
			keyName = keyName[:len(keyName)-len(filepath.Ext(key.Name()))]

			nameParts := strings.Split(keyName, ".")
			keyName = nameParts[0]
			keyType := nameParts[1]

			f, err := os.Open(fmt.Sprintf("data/redis/%s/%s", app.Name(), key.Name()))
			if err != nil {
				return errors.Trace(err)
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				return errors.Trace(err)
			}
			strContent := string(content)

			data := map[string]interface{}{}
			if err := hcl.Decode(&data, strContent); err != nil {
				return errors.Trace(err)
			}

			if len(data) != 1 {
				return errors.New("there is more than one root key, one key per file only")
			}

			for protoName, values := range data {
				entities := values.([]map[string]interface{})

				log.WithFields(log.Fields{
					"app":      app.Name(),
					"key-name": keyName,
					"entities": len(entities),
				}).Info("Fill data")

				for _, entity := range entities {
					var buf bytes.Buffer
					if err := json.NewEncoder(&buf).Encode(entity); err != nil {
						return errors.Trace(err)
					}

					m := jsonpb.Unmarshaler{}
					msg := reflect.New(proto.MessageType(protoName).Elem()).Interface().(proto.Message)
					if err := m.Unmarshal(bytes.NewReader(buf.Bytes()), msg); err != nil {
						return errors.Trace(err)
					}

					switch keyType {
					case "queue":
						if err := sess.Queue(keyName).Push(msg); err != nil {
							return errors.Trace(err)
						}

					default:
						return errors.Errorf("unexpected key type: %s", keyType)
					}
				}
			}
		}
	}

	return nil
}
