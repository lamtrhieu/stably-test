# Stably Prime Test

I have try to use golang to implement a simple html web app. 

## Usage 

- Clone the code from `git@github.com:lamtrhieu/stably-test.git`
- Build the docker file <code>docker build -t myapp . </code>
- Check the image is built successfully with  <code>docker images</code>
- Run the image <code>docker run -p 8080:8081 -it myapp</code>
- Access http://localhost/

## Approach 
- To solve the problem, I just use a simple loop to decrease from the input until the index number is a prime. 
- A simple check prime number is implemented
- A cache is provided to speed up perf


# Know limitation 
- Cache is a in-memory map and is reset whenever server restart. Should use Redis or memcache in production. But for the exercise, in-memory cache is used to demonstrate the idea.
- Have not use any advance feature of golang like module, package, dependency injection library like wire 
- Still make use of global variable in main. Should use dependency injection library and abstract that away instead. 
- The app is simple client/server app and require page reload for every request. Should use SPA like React or Angular so that html client is more responsive. 
- Validation should be added at client side. For now, it's only implemented at server side. 



