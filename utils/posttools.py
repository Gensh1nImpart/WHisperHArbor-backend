import http.client
import json

conn = http.client.HTTPConnection("localhost", 8000)

headers = {
  'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImFjY291bnQiOiJ5cmgiLCJwYXNzd2QiOiIxMjMzMjEifSwiZXhwIjoxNjk5NTQ3NDg1LCJpc3MiOiJ5cmgifQ.hgLl1GCUpZEEbL7EG176DHL5e-Juf3DcdtVXehhnk9U',
  'Content-Type': 'application/json'
}
for i in range(15, 50):
  payload = json.dumps({
    "content": f"test${i}",
    "encrypted": False
  })
  conn.request("POST", "/v1/api/post", payload, headers)
  res = conn.getresponse()
  data = res.read()
  print(data.decode("utf-8"))