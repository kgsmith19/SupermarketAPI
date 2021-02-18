This is a Supermarket API excerise created by Kyle Smith, programmed in Golang.  This application was designed to fetch, add, and delete supermarket inventory items from a database.  

Notes: 
-Adding multiple inventory entries at once will not successfully post.
-There is no Docker image within this program.

Assumptions made:
-API testers have basic knowledge with REST & SOAP API request testing tools.
-"$" is not a valid character to enter a unit price. 
-Produce Item Name can be any non-empty(empty includes only blank spaces) string.
-Unit price can be entered as an integer or decimal.

Application execution instructions:
-After reaching the repository location(https://github.com/kgsmith19/SupermarketAPI), download the code from master branch.
-Within the main folder run SupermarketAPI.exe(This will run in the background).

-Go to http://localhost:10000 with your internet browser.
-You should see a message welcoming you to the Grocery Store Inventory API.
-Go to your favorite REST testing tool.  For example, reqbin.com

--GET--
--fetch all inventory
http://localhost:1000/inventory <--This will fetch all inventory to the database.  Set GET as request type if using testing tool.

--fetch produce item
http://localhost:1000/produceItem/{produceCode} <--This will fetch a single item if produce code is found in the database. Set GET as request type if using testing tool.
e.g. http://localhost:1000/produceItem/A12T-4GH7-QPL9-3N4M

--POST--
--add produce item
http://localhost:10000/addProduceItem <-- This will add a single produce item to the database.
Go to REST Testing tool and enter url with POST desigination.  There should be another location to enter JSON Content for the produce item to be added. 
e.g. {"ProduceCode":"XUIT-212D-LDSW-256F","":"Orange","UnitPrice":"5.67"}

--delete produce item
http://localhost:1000/deleteProduceItem/{produceCode} <--This will delete a single produce item if found in the database.  Set POST as request type.
e.g. http://localhost:1000/deleteProduceItem/A12T-4GH7-QPL9-3N4M














