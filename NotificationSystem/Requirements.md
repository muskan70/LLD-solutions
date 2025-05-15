## Problem Statement: Design a Notification Scheduler Service

You need to implement a service that schedules and sends notifications (email, SMS, or push) to users at a specific time. The system should allow creating, updating, deleting, and listing scheduled notifications. For this round, assume actual sending is mocked.

### Requirements:
- Create Notification – Schedule a notification with:
> UserID
> Channel (Email, SMS, Push)
> Message
> ScheduledTime (UTC)
> Priority (LOW, MEDIUM, HIGH)
- Update Notification – Update the scheduled time or message of a pending notification.
- Delete Notification – Cancel a scheduled notification.
- List Notifications – List all scheduled notifications for a user.
- Notification Dispatcher – Periodically check for due notifications and "send" (mock send) them.

### Bonus Extensions :
- Add retry logic for failed sends.
- Add a filter in ListNotifications by status.
- Simulate storage using a local JSON file (instead of in-memory).
- Add basic unit test(s).



