First thing is that it only supports a post Request
with the following params
````
{
    "val":"any_string"
}

{
    resize:n
}
here n should be an integer
````
About the implementation there are 3 versions<br/>
1.In master is a simple pool which uses sync waitgroup
and a very simple worker.</br>

2.In v2 branch is a simple dispatcher model
  There is a small bug here on closing channels ;( </br>

3.In v3 branch is another implementation of the same pool but
  with a little more elegant solution.  
 