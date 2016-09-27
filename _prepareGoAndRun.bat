@echo off 
call _batch\typescriptCompile.bat
call _batch\go-bindata.bat
call _batch\gorazorExec.bat

go run server.go