# Design Patterns in go

## Creational design pattern
    - Abstract design Pattern
    - Singleton Pattern
    - Builder
    - Factory
    - Object Pool
    - Prototype


### Abstract Factory Design Pattern in Go

#### Definition:

  Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects. It is an abstraction over the factory pattern.

#### Example use cases:

Imagine you need to buy a sports kit which has a shoe and short. Preferably most of the time you would want to buy a full sports kit of a similar factory i.e either nike or adidas. This is where the abstract factory comes into the picture as concrete products that you want is shoe and a short and these products will be created by the abstract factory of nike and adidas.
Both these two factories – nike and adidas implement iSportsFactory interface.
We have two product interfaces.

iShoe – this interface is implemented by nikeShoe and adidasShoe concrete product.
iShort – this interface is implemented by nikeShort and adidasShort concrete product.

- [Note](https://golangbyexample.com/abstract-factory-design-pattern-go/)

### Builder Pattern in GoLang

Builder Pattern is a creational design pattern used for constructing complex objects.

When To Use:

- Use Builder pattern when the object constructed is big and requires multiple steps. It helps in less size of the constructor.  The construction of the house becomes simple and it does not require a large constructor

- When a different version of the same product needs to be created. For example, in the below code we see a different version of house ie. igloo and the normal house being constructed by iglooBuilder and normalBuilder

- When half constructed final object should not exist. Again referring to below code the house created will either be created fully or not created at all. The Concrete Builder struct holds the temporary state of house object being created

### Factory Design Pattern in Go

Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern. This pattern provides a way to hide the creation logic of the instances being created.
The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

When To Use:

- We have iGun interface which defines all methods a gun should have
- There is gun struct that implements the iGun interface.
- Two concrete guns ak47 and maverick. Both embed gun struct and hence also indirectly implement all methods of iGun and hence are of iGun type
- We have a gunFactory struct which creates the gun of type ak47 or maverick.
- The main.go acts as a client and instead of directly interacting with ak47 or maverick, it relies on gunFactory to create instances of ak47 and maverick

### Object Pool Design Pattern in Go

The Object Pool Design Pattern is a creational design pattern in which a pool of objects is initialized and created beforehand and kept in a pool. As and when needed, a client can request an object from the pool, use it, and return it to the pool. The object in the pool is never destroyed.

When to Use:
- When the cost to create the object of the class is high and the number of such objects that will be needed at a particular time is not much.
    - Let’s take the example of DB connections. Each of the connection object creation is cost is high as there is network calls involved and also at a time not more than a certain connection might be needed. The object pool design pattern is perfectly suitable for such cases.

- When the pool object is the immutable object
    - Again take the example of DB connection again. A DB connection is an immutable object. Almost none of its property needs to be changed

- For performance reasons. It will boost the application performance significantly since the pool is already created

### Prototype Pattern in Go

It is a creational design pattern that lets you create copies of objects. In this pattern, the responsibility of creating the clone objects is delegated to the actual object to clone.

The object to be cloned exposes a clone method which returns a clone copy of the object

When to Use
- We use prototype pattern when the object to be cloned creation process is complex i.e the cloning may involve vases of handling deep copies, hierarchical copies, etc. Moreover, there may be some private members too which cannot be directly accessed.
- A copy of the object is created instead of creating a new instance from scratch. This prevents costly operations involved while creating a new object such as database operation.
- When you want to create a copy of a new object, but it is only available to you as an interface. Hence you cannot directly create copies of that object.

### Singleton Pattern in Go

Singleton Design Pattern is a creational design pattern and also one of the most commonly used design pattern. This pattern is used when only a single instance of the struct should exist. This single instance is called a singleton object. Some of the cases where the singleton object is applicable:

- DB instance – we only want to create only one instance of DB object and that instance will be used throughout the application. 
- Logger instance – again only one instance of the logger should be created and it should be used throughout the application.
