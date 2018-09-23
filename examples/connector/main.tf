provider "kafkaconnect" {
  url = "https://kafkaconnect.example.com"
}

resource "kafkaconnect_connector" "file_stream_source" {
  name = "local-file-source"
  class = "FileStreamSource"
  maximum_tasks = "1"

  configuration = {
    file = "/tmp/test.txt"
    topic = "connect-test"
  }
}
