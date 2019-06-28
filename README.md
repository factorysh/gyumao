# Gyūmaō

[![Build Status](https://drone.bearstech.com/api/badges/factorysh/gyumao/status.svg)](https://drone.bearstech.com/factorysh/)

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
