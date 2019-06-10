# Gyūmaō

Bird view

```
    telegraf
        |
        v
      relay -> influxdb
        |
        v
     gyumao
```

point has hierarchies :

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
