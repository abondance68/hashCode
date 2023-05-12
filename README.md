 # Hashcode <br>
 
 The user will be prompted in the CLI to enter his name or any name he/she chooses.<br>
 Each letter of the name will then be converted into its ASCII code.<br>
 Once each letter is transformed into its number in ASCII , all the generated numbers are added and this  sum is then divided by 100.<br>
 We chose 100 because it is the size of the array in which we will store the name into in the case of a hash table.<br>
 The resulting hash will always be , in this case, an integer between 0 & 99.<br> 
 
This hashing algorithm is the prerequisite for building Hash Tables.<br>
Nevertheless, if two names , after being hashed with the hashing algorithm, results in the same output, then we have<br>
a condition called  'collision'  and we need a way to handle collision.<br>

There exist 2 ways for handling collisions, one of the most reliable & effective is 'separate chaining' . <br>

'Separate chaining' means we store multiple names into one slot/index of the array, and <br>
we do that using linked lists.<br>
(in our case, the size of the array has been chosen to be '100' but it can be any number) . <br>


One of the advantages of linked lists is that such data structures will always be able to grow and shrink dynamically, <br>
thus, we do not have much limitations for the size.<br>
In a nutshell, each slot/index of our array will hold a pointer that points to the head of a linked list<br> 
tha has a list of names.<br>







  

