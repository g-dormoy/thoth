# thoth

Is a service mesh

## Add a service

```bash
curl -X POST http://localhost:4242/add -H "Content-Type: application/json" --data '{"name": "service1", .....}'
```

## Get a service

```bash
curl http://localhost:4242/get?name=service1
```
