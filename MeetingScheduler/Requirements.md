## Design a Meeting Scheduler to book a room.

### Problem Statement:
1. Given N rooms and M users. A host(user) provides list of participants and duration of meeting. Assign an available slot and vacant room when all users are available.  

### Requirements:
1. Check all available slots of given duration of all participants to schedule a meeting.
2. Check the available room of the available slots and assign that room.
3. Notify all users with all details related to meeting i.e date, time, roomId, meeting agenda, content related to meeting.
4. User can check history of his meeting scheduled, Also room scheduled meeting can be checked.
5. Different room assignment strategies can be present.
6. Capacity of room and type of meeting room can to be defined.
7. Also meeting can be reccurent or one time?


### Core Components:
1. User
2. Room
3. Calender
4. Meeting
5. Meeting Scheduler

### Interfaces
1. MeetingObservable: NotifyUsers()
2. Observer : Notification() 
3. RoomAllocationStrategy : BookRoom()


### Class Definations
1. Meeting 
- Attributes => MeetingId, Date, day, StartTime, EndTime, RoomId, Host: User, Participants: List<User>, MeetingType: Recurrent/ OneTime
- Methods    => AddUser(), NotifyUsers(), UpdateSchedule(), SetMeetingRecurrentSchedule()

2. Calender: 
- Attributes => map<date, List<MeetingId>>
- Methods    => GetMeetingsForDay(), IsAvailable(), GetAvailableSlots(), GetScheduledMeetingsForDay(), ScheduleMeeting()

3. User extends Calender
- Attributes => UserId, name, email, phone, UserCalender: Calender
- Methods    => All Calender Methods

4. Room extends Calender
- Attributes => RoomId, Capacity, RoomType: teemMeet/seminar, RoomCalender: Calender
- Methods    => CheckRoomPreference(), All Calender Methods

5. RoomAllocationStrategy : FirstComeFirstServeStrategy , PriorityBasedStrategy

### Design Patterns
1. Strategy Design Pattern for room allocation
2. Observer Design Pattern to notify users




