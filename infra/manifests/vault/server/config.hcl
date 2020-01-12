ui = true
disable_mlock = true

listener "tcp" {
  tls_disable = true
  address = "[::]:8200"
  cluster_address = "[::]:8201"
}

storage "raft" {
  path = "/raft"
}
