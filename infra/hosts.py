#!/usr/bin/env python

# Add here your hosts...
hosts = [
  'admin.app',
  'admin-assets.app',

  'api.app', 
]


lines = []
with open('/etc/hosts', 'r') as f:
  for line in f.readlines():
    if '.app' in line:
      continue
    lines.append(line)


with open('/etc/hosts', 'w') as f:
  f.write(''.join(lines))
  for host in hosts:
    f.write('127.0.0.1\t%s\n' % host)
