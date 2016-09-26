dir content\js\ts\*.ts /b /s > ts-files.txt
call tsc --removeComments --outDir content\js\ @ts-files.txt
del ts-files.txt