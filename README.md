# graphFinder

- This is just a tool to search for GraphQL endpoints. It also check if introspection is enabled.


### Obs
- This is my first goolang tool, you may find a lot of bugs.
- I know I should change/add criterias to confirm the GraphQL existence.
- I'll appreciate any hints about the code or the tool.


### Usage 
```

   ____     ____        _       ____    _   _    _____              _   _    ____  U _____ u   ____     
U /"___|uU |  _"\ u U  /"\  u U|  _"\ u|'| |'|  |" ___|    ___     | \ |"|  |  _"\ \| ___"|/U |  _"\ u  
\| |  _ / \| |_) |/  \/ _ \/  \| |_) |/| |_| |\U| |_  u   |_"_|   <|  \| |>/| | | | |  _|"   \| |_) |/  
 | |_| |   |  _ <    / ___ \   |  __/ U|  _  |u\|  _|/     | |    U| |\  |uU| |_| |\| |___    |  _ <    
  \____|   |_| \_\  /_/   \_\  |_|     |_| |_|  |_|      U/| |\u   |_| \_|  |____/ u|_____|   |_| \_\   
  _)(|_    //   \\_  \\    >>  ||>>_   //   \\  )(\\,-.-,_|___|_,-.||   \\,-.|||_   <<   >>   //   \\_  
 (__)__)  (__)  (__)(__)  (__)(__)__) (_") ("_)(__)(_/ \_)-' '-(_/ (_")  (_/(__)_) (__) (__) (__)  (__) 


  -f string
        Input File name.
  -o string
        Output File name.
  -proxy string
        Set Proxy (ex: http://localhost:8080)

```

### Output
```
[+] [FOUND] - https://rickandmortyapi.com/graphql
[+] [FOUND] - https://countries.trevorblades.com/console/graphql
[+] [FOUND] - https://countries.trevorblades.com/qql
[+] [FOUND] - https://countries.trevorblades.com/graphql
[+] [FOUND] - https://countries.trevorblades.com/api/graphql
[+] [FOUND] - https://countries.trevorblades.com/graphiql
[+] [FOUND] - https://countries.trevorblades.com/ql
[+] [INSTROSPECTION ENABLED] - https://rickandmortyapi.com/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/console/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/qql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/api/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/graphiql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/ql

```

### Examples
```
cat targets.txt | graphFinder -o outputfile.txt
```
```
graphFinder -f targets.txt -o outputfile.txt -proxy http://localhost:8080
```

```
cat targets.txt | graphFinder -o outfile.txt


   ____     ____        _       ____    _   _    _____              _   _    ____  U _____ u   ____     
U /"___|uU |  _"\ u U  /"\  u U|  _"\ u|'| |'|  |" ___|    ___     | \ |"|  |  _"\ \| ___"|/U |  _"\ u  
\| |  _ / \| |_) |/  \/ _ \/  \| |_) |/| |_| |\U| |_  u   |_"_|   <|  \| |>/| | | | |  _|"   \| |_) |/  
 | |_| |   |  _ <    / ___ \   |  __/ U|  _  |u\|  _|/     | |    U| |\  |uU| |_| |\| |___    |  _ <    
  \____|   |_| \_\  /_/   \_\  |_|     |_| |_|  |_|      U/| |\u   |_| \_|  |____/ u|_____|   |_| \_\   
  _)(|_    //   \\_  \\    >>  ||>>_   //   \\  )(\\,-.-,_|___|_,-.||   \\,-.|||_   <<   >>   //   \\_  
 (__)__)  (__)  (__)(__)  (__)(__)__) (_") ("_)(__)(_/ \_)-' '-(_/ (_")  (_/(__)_) (__) (__) (__)  (__) 

 USAGE: ./GraphFinder -f inputfile.txt -o outputfile.txt


[+] [FOUND] - https://rickandmortyapi.com/graphql
[+] [FOUND] - https://countries.trevorblades.com/ql
[+] [FOUND] - https://countries.trevorblades.com/graphiql
[+] [FOUND] - https://countries.trevorblades.com/api/graphql
[+] [FOUND] - https://countries.trevorblades.com/qql
[+] [FOUND] - https://countries.trevorblades.com/console/graphql
[+] [FOUND] - https://countries.trevorblades.com/graphql
[+] [INSTROSPECTION ENABLED] - https://rickandmortyapi.com/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/ql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/graphiql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/api/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/qql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/console/graphql
[+] [INSTROSPECTION ENABLED] - https://countries.trevorblades.com/graphql
```