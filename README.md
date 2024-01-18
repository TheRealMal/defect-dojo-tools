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
|-:|:-:|:-:|
|create_product||Creates new DefectDojo product with provided details|
||name||
||description||
||tags||
||product_type||
||product_sla||
|find_product||Finds DefectDojo product by provided product name|
||name||
|create_engagement||Creates new DefectDojo engagement with provided details|
||product_id||
||name||
||description||
||commit_hash||
||branch_tag||
||status|`Not Started`/`Blocked`/`Cancelled`/`Completed`/`In Progress`/`On Hold`/`Waiting for Resource`|
|upload_report||Uploads report to provided engagement id|
||engagement_id||
||report_format||
||report_filename||
||close_old_findings|`true` or `false`|