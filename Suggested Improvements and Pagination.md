# Suggested Improvements for the Get All Folders Component
After going through  the `GetAllFolders` function, I made some improvements to ensure that the program could run correctly, and was also made more efficient.

The improvements are:
1. Removing the f1 Folder from the defined variables at the start of the function as it was not required.
2. Removing the `f := []Folder{}` line and instead, using the `fs` list of folders defined at the start of the function.
3. Utilising the `err` variable when calling the `FetchAllFoldersByOrgID` in case any errors are returned. An if statement was also added to check if an error is returned, so that the function can stop executing and return the error.
4. Changing the for loop over the range of returned folders to the line `fs = append(fs, r...)`. The use of the `r...` allows each element returned in the `r` slice to be appended to `fs`. This simplifies the code as the for loop can be removed.
5. Remove the line `var ffr *FetchFolderResponse` as this declaration can be merged with the assignment line instead: `ffr := &FetchFolderResponse{Folders: fs}`.

With these improvements, the code was able to run successfully, and return a correct json body of all folders with the provided Organisation Id.


# Pagination Solution
The pagination solution that I developed involved the addition of a few token and index variables that would allow the pagination to occur within the context of the program.

This solution was implemented with a number of changes:
1. A new struct, `FetchFolderResponsePaginated` was created in the `types.go` file. This was similar to the `FetchFolderResponse` struct, with the addition of a `NextToken` field. This `NextToken` field is a string, and would refer to the id of the next folder that would be returned in the next page of results. This field is returned in each response in the JSON body, which can be copied and added to the next function call to retrieve the next page of results. 
2. A `pageLimit` constant was added and set to 10. This constant would ensure that each page of results would not exceed 10 folders. This constant can be changed to any number.
3. The addition of a `token` string argument to both functions. This token would either be an empty string, indicating that the requested response will be the first page of results, or, it would be the `NextToken` field provided in a previous response body. This token is used as the starting point of the folder list that is returned.
4. The `FetchAllFoldersByOrgIDPaginated` function involved the below changes:
    - Setting the start position of the fetch to be the id of the folder provided by the token argument.
    - Looping through the folders from the start position and only adding a maximum of 10 folders to the response.
    - Generating the next token by finding the next folder. If no folder is found, the next token will be empty, indicating the final page of the response has been returned.