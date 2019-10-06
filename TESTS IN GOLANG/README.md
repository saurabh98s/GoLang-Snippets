# TESTING WITH GOLANG
<br>
Testing is hugely important in all software. Being able to ensure the correctness of your code and ensure that any changes you make don’t end up breaking anything else in different parts of your codebase is hugely important.

By taking the time to adequately test your go programs you allow yourself to develop faster with a greater sense of confidence that what you are developing will continue to work when you release it to production.
<br>
## Have a look on a quick tutorial by visiting the link below :
(https://youtu.be/GlA57dHa5Rg)

## This is how you can actually view how much of your code is covered by tests
With the important part out of the way, let’s look at how you can check the test coverage of your system using the go test command:

Within the same directory as your main.go and your main_test.go files, run the following:<br>
    **go test -cover**
<br>
to further visualise the code like in the screenshot, you can just write:<br>
    **go tool cover -html=coverage.out**
<br>



