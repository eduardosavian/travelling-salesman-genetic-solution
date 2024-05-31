# Knight's tour algorithms

[Vídeo](https://youtu.be/rsfa4lfuNqg)

## Description

Implementar a solução para o problema do Caixeiro Viajante utilizando Algoritmos Genéticos.

O programa deve permitir que o usuário defina a quantidade de cidades e a distância entre elas. O resultado deve ser a melhor rota possível, mostrando a sequencia de cidades a ser precorrida e a distância total da rota.

Entrada: Ler de arquivo texto, com ; (ponto e virgula) para separar os dados, conforme exemplo que segue:

5; 10; 15; 5; 7; 20; 9; 11; 10; 18; 19

4; 2; 5; 10; 8; 4; 7

- Onde o primeiro valor é a quantidade de cidades e os demais a distancia entre elas. No primeiro exemplo são 5 cidades, da cidade 1 para 2 a distância é 10, de 1 para 3 é 15 e assim sucessivamente (conforme tabela abaixo). No segundo problema a quantidade de cidades é 4.

- o número máximo de cidades é de 20 e o arquivo pode ter N problemas a serem resolvidos. No exemplo de entrada o N é 2.

Obs.: O trabalho deve ser apresentado pela equipe no dia da entrega.

## Run

### Windows

#### Go build

```ps1
go build -o libs/knight_tour src/main.go src/methods.go
```

#### Create Python virtual environment

```s1
python -m venv venv
```

#### Activate Python virtual environment

```ps1
.\venv\Scripts\activate
```

#### Install Python requirements

```ps1
pip install -r requirements.txt
```

#### Create Python executable

```ps1
pyinstaller --onefile --name knights_tour --distpath exec --add-data "libs:libs" --hidden-import seaborn src/menu.py
```

#### Move executable

```ps1
mv exec/knights_tour.exe knights_tour.exe

```

#### Run executable

```ps1
.\knights_tour.exe
```

### Linux

#### Go build

```bash
go build -o libs/knight_tour src/main.go src/methods.go
```

#### Create Python virtual environment

```bash
python -m venv venv
```

#### Activate Python virtual environment

```bash
source venv/bin/activate
```

#### Install Python requirements

```bash
pip install -r requirements.txt
```

#### Create Python executable

```bash
pyinstaller --onefile --name knights_tour --distpath exec --add-data "libs:libs" --hidden-import seaborn src/menu.py
```

#### Move executable

```bash
mv exec/knights_tour knights_tour
```

#### Run executable

```bash
./knights_tour
```

## Bibliographic references

[An efficient algorithm for the Knight’s tour problem - Ian Parberry](https://core.ac.uk/download/pdf/81964499.pdf)

[Knight's tour - Wikipedia](https://en.wikipedia.org/wiki/Knight's_tour)
