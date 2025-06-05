## Implement a simplified version of a banking system.

### Level 1:
**Ask**: Please implement a banking system library with methods that can support new account creation as well as depositing and withdrawal functions. <br>
**Operations**:
1. deposit <accountId> <amount>
> Should create a new account with the given identifier and starting balance if it doesn't already exist.<br>
> If the account already exists, deposit the given `amount` of money to the specified account `accountId`<br>
> Returns "true" if an account was successfully created, <br>
> Returns `amount` of the specified account if an account with `accountId` already exists. <br>

2. withdraw <accountId> <amount>
> Should withdraw the given `amount` of money to the specified account `accountId` <br>
> Returns a string representing the total amount of money in the account after the query has been processed. <br>
> If the specified balance of the account is lower than the `amount` to be withdrawn, it should return an empty string. <br>
> If the specified account doesn't exist, it should return an empty string <br>

### Level 2:
**Ask**: On your existing implementation, please implement a new functionality that would support the transferring of funds from one account to another.<br>
**Operation**: transfer <fromId> <toId> <amount>
> Should transfer the given amount of money from account with fromId to account with toId. <br>
> Returns a string representing the balance of fromId if the transfer was successful, or an empty string otherwise. <br>
> Returns an empty string if fromId or toId doesn't exist.<br>
> Returns an empty string if fromId and toId are the same.<br>
> Returns an empty string if funds on the account fromId are insufficient to perform the transfer.<br>


### Level 3:
**Ask**: On your existing implementation, please implement a new functionality that would allow a user to identify the top N accounts sorted by spending activity (including transfers out and withdrawals). <br>
**Operation**: TOP_SPENDERS <n>
> Should return identifiers of top n accounts sorted by the total amount of money moved out of the account - either transferred out or withdrawn, in descending order. In case of a tie, sorted alphabetically by accountId in ascending order. 
> The result should be a string in the following format: "accountId_1 (totalMovedOut_1"), "accountId_2 (totalMovedOut_2)" , ..., "accountId_n (totalMovedOut_n)".
> If less than n accounts exist in the system, then return all their identifiers (in the described format).

Example:
The example below shows how these operations should work:<br>
operations = [<br>
  ["DEPOSIT", "account1", "1000"],<br>
  ["DEPOSIT", "account1", "500"],<br>
  ["DEPOSIT", "account2", "1000"],<br>
  ["WITHDRAW", "non-existing", "2700"]<br>,
  ["WITHDRAW", "account1", "2000"],<br>
  ["WITHDRAW", "account1", "500"],<br>
  ["TRANSFER", "account1", "account2","1001"],<br>
  ["TRANSFER", "account1", "account2",<br> "200"]<br>
  ["TOP_SPENDERS", "2"]<br>
  ["WITHDRAW", "account2", "800"],<br>
  ["TOP_SPENDERS", "3"],<br>
  ["WITHDRAW", "account1", "100"],<br>
  ["TOP_SPENDERS", "2"]<br>
]<br>
<br>
returns "1000"<br>
returns "1500"; an account with this identifier already exists<br>
returns "1000";<br>
returns ""; an account with this identifier doesn't exist<br>
returns ""; withdrawal amount exceeds the account balance<br>
returns "1000"<br>
returns ""; this account has insufficient funds<br>
returns "800"<br>
returns "account1(700), account2(0)"<br>
returns "400"<br>
returns "account2(800), account1(700)"<br>
returns "700"<br>
returns "account1(800), account2(800)"<br>
