curl -X POST http://localhost:8000/notes \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Fixed Create command after setting up all the layers ",
    "content": "Learning Go, Gin, MongoDB step by step"
  }'