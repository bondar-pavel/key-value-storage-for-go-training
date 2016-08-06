# key-value-storage-for-go-training
Key-value storage written on Go just for training purposes.

# Task
You are encouraged to implement the key-value storage server.
Keys and values are utf8-encoded strings.

SET method updates one key at a time with the given value.

GET method returns tuple of the value and the key state. The state either
present or absent.

DEL method removes one key at a time and returns the state of the resource.
The state either ignored or absent. A key should be ignored if it does not exist

# Example of the CLI (Telnet) interface
```
> SET 4:key1 5:Hello 
> SET 4:key2 5:World 
>
> GET 4:key1 
> 5:Hello (present) 
>
> GET 4:key3 
> 0: (absent)
>
> // Added by Pavel
> DEL 4:key1
> 1: (deleted)
>
> DEL 4:key3
> 0: (absent)

```
