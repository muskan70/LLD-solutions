## Chat Server Application

### Requirements
1. System should be able to support real time chat rooms for groups as well as direct messaging.
2. A group can consist of max 10 people.

### Core Components
1. Chat
2. User
3. Group
4. Message
5. Chat System

### Design Patterns
1. Mediator Pattern : Chat system act as mediater between users to communicate.
2. Publish Subscribe Pattern : Messages should be published by sender so that other user or group participants can consume them.