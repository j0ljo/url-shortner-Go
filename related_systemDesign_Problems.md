# System's Design Problems 

Main parts:
1. Functional Requirements 
2. Non-Functional Requirements

### 1. Functional Requirements 

```
i) Given a Long url ---> Create Short url 
ii) Given a short url ---> Redirect to long url 
```

### 2. Non-Functional Requirements 
```

i) Low latency 
ii) High availability 

```

```
// Clarifying Questions: 

Q1. How many url's will be created per second? 
- 1000 url's per second 

Q2. What character can we use? 
- alphanumeric (A-Z, a-z, 0-9) = 62 characters 

Q3. What would happen if multiple users have the exact same long url? 
- Always create a new short url for each user 
```
```

```

## Data Estimation 

Q. Number of unique url's needed (10 years) 
- seconds in a year : 60*60*24*365 = 31.5M 
  total seconds in 10 years : 31.5M * 10 = 315M 
  total url's in 10 years : 1000 * 315M = 315B 

Q. Length of url identifier
 
- 62^1 = 62 unique urls (1 character) 
  62^2 = 3644 unique urls (2 characters) 
  .
  . 
  62^7 = 3.5T unique urls (7 characters)

Q. How much data storage would we need? 
```
```

```
  short url     -> 7 bytes 
  long url      -> 100 bytes 
  user metadata -> 500 bytes 
  total         -> ~1000 bytes 

  Approximately 1000 byte per url ,SO 
  1000*total no. of urls for 10 years ( 315M) = 315 Terabytes of data 
```


## High level system design 

## API design 
