# defect-dojo-tools
Tools for easier and better DefectDojo experience

## Guide
First you run tool with necessary command with DefectDojo API URL & Token
```console
./ddtool <COMMAND> <API_URL> <API_TOKEN>
```
Then you write each value for input from a new line  

Example:
```console
./ddtool create_product http://localhost:8080 qwerty12345
Example New Product
Some description
Tag 1,Tag 2,Tag 3
1
1
```
## Commands list
|Command|Inputs|Explanation|
|-|-|-|
|create_product||Creates new DefectDojo product with provided details|
||name||
||description||
||tags||
||product_type||
||product_SLA||
|find_product||Finds DefectDojo product by provided product name|
||name||
