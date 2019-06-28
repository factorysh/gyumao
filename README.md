# Gyūmaō

![Build Status](https://drone.bearstech.com/api/badges/factorysh/gyumao/status.svg)](https://drone.bearstech.com/factorysh/)

[牛魔王 Ox demon king](https://en.wikipedia.org/wiki/List_of_Dragon_Ball_characters#Ox-King)

Get measurements in _influxdb_ format, apply rules, send alert to _alertmanager_.

## Bird view

Short arm:

```
telegraf
    |
    v
 gyumao -> alertmanager
```

Long arm:

```
telegraf
    |
    v
  relay -> influxdb
    |
    v
 gyumao -> alertmanager
```

Point has hierarchies :

- DC / server
- Client / project / service

point has a key : name + list of tags

```
    http
      | points
      v
    chan
      | points
      v
    filter on point name
      | point
      v
    rules on a go routine

```

last time seen.

circular store, ordered by date.

## Licence

GPL v3. © 2019 Mathieu Lecarme.
