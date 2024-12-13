

# go run -exec "goexec" "http.ListenAndServe(':8080', http.FileServer(http.Dir('.')))"


echo
echo "launching server at   http://localhost:8080/"
echo

go run local_server.go


