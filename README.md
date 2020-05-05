# Библиотека proto файлов и их компиляции

> собрать образ:   
`bash build.sh`


> запустить генерацию классов:   
`bash run.sh`   


### pkg
сгенерированные классы здесь: `pkg/0.0.0/{python,golang,java}`



### ссылки
<pre>
python
https://grpc.io/docs/quickstart/python/









</pre>




-----

подготовка плагинов для protoc
git clone --recursive https://github.com/grpc/grpc
cd grpc && make plugins -j 12
ls bins/opt/grpc_python_plugin

# тоже, только один плагин для python
make grpc_python_plugin
