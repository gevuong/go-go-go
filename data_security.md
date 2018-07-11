# Intro to Data Security
- Keeping data and communications secure is one of the most important topics in development today.

## Things to consider when thinking about Data Security, Risk Assessment and Management
- Why would someone want to gain access to your application or data?
- What are the most likely ways someone would gain that access?
- What data would be the most valuable to an attacker?
- Assuming your application is compromised, how would that impact your users?
- How would it impact your business?
- What processes would you follow to fully recover from an attack?
- How could someone use your application for other illicit uses.

## Hash
- A one-way street, meaning data that has been hashed, cannot be "unhashed". 
- Hash collisions occur when two different inputs generate the same hash. Hash collisions are what compromises hashing algorithms.
- To prevent hash collisions, add a salt, which is additional data, to the hash. For example, a salt can be the timestamp in nanoseconds when a password was created or when a user signed up.
- MD5 algorithm was once a very commonly used hashing function in the earlier days. It's now severely compromised because it doesn't require large computing power to find hash collisions between two inputs.
- But what if we want to safely store data and be able to read it? Lets say you want to store sensitive data in the database, and save it for a later time. You would want to store the encrypted version in the database, because it's a two way street. This is where encryption comes into play.

## Encryption
- Is a two-way street, and is most commonly used with communication.
- Uses an algorithm to transform information into unreadable format, and requires a "key" to decrypt data into its original.
- Turns readable data (i.e. plain text) into unreadable, encrypted data (i.e. cipher text), and vice versa.
- Examples: 
    - Connecting to a website that uses an SSL (Secure Sockets Layer) certificate (i.e. HTTPS), the communication you and the website is encrypted, and is thus secret. SSL certificate encrypts data being sent
    - Without HTTPS, when you submit a form that has your password or username, anyone that is tapped to the connection between you and the server can see the data. Having an SSL certificate encrypts all the data so no one can inspect because it's unreadable.
    - Messages sent from WhatsApp or Signal are also encrypted.

### When would you want to use Encryption?
- encrypt data you're storing for legal purposes. Callisto is an app that allows users to report sexual assaults on college campuses. Callisto encrypts all reports before storing them in DB so that only the person who reported it can securely access them. 


- There are two main styles of encryption.

### Symmetric vs Asymmetric (public key) encryption
**Symmetric (or private key encryption)**
- plain text is encoded and decoded using the same private key. A single private key is used among all parties.
- This means that if private key is ever compromised, data can be decrypted, stolen, and abused.
- A fast process since it uses a single key.
- Appears to be less popular nowadays, but blowfish, AES, and Idea all use symmetric encryption.

**Asymmetric (or public key encryption)**
- Everyone has two keys, a public and a private. The public key is distributed, while the private key is never shared. This means that anyone can encrypt messages using the public key, but only the holder of the paired private key can decrypt the message.
- For example, if you want to send me an encrypted message, you would encrypt it using my public key. And I would decrypt it using my private key.
- For example, if Alice and Bob are using asymmetric encryption to chat with each other, then there are a total of 4 keys.
- In private key encryption, Alice would use Bob's public key to encrypt a message for him.

- SSL and SSH both use a combination of public key and private keys.

## Additional Resources
- Doing Login Systems Wrong `https://ticki.github.io/blog/you-are-probably-doing-login-systems-wrong/`
- SHA-1 broken by Google's Project Zero/Google Research Security `https://shattered.io/`
