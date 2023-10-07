# Composition over Inheritance

## Problems with Inheritance
- Large hierarchies
- Tight coupling

## What is Inheritance?
Is-a relationship
```
Animal - Eat(), Sleep() + Walk()
  |_ Person - Eat(), Sleep() + Walk()
  |_ Dog - Eat(), Sleep() + Walk()?
  |_ Goldfish - Eat(), Sleep() + Walk()!
```
```
Animal - Eat(), Sleep()
  |_ Mammal - Walk()
     |_ Person - Eat(), Sleep(), Walk()
     |_ Dog - Eat(), Sleep(), Walk()
  |_ Fish - Swim()
     |_ Goldfish - Eat(), Sleep(), Swim()
```

## What is Composition?
Has-a relationship
```
Alive - Eat(), Sleep()
Walkable - Walk()
Swimmable - Swim()
```

```
Person 
   |_ Alive - Eat(), Sleep()
   |_ Walkable - Walk()
Dog 
   |_ Alive - Eat(), Sleep()
   |_ Walkable - Walk()
Goldfish
   |_ Alive - Eat(), Sleep()
   |_ Swimmable - Swim()
Duck
   |_ Alive - Eat(), Sleep()
   |_ Walkable - Walk()
   |_ Swimmable - Swim()
```