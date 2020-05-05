### плагин для python нужно собирать самостоятельно

<pre>
автоматизации на это нет. Плагин собран 05/05/2020 
часто не меняется. можно смотреть на исходный репозиторий 


подготовка плагинов для protoc
git clone --recursive https://github.com/grpc/grpc
cd grpc && make plugins -j 12
ls bins/opt/grpc_python_plugin

# тоже, только один плагин для python
make grpc_python_plugin

</pre>
