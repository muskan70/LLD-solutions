## Design and implement a Job Scheduler

### Requirements
1. Accept jobs submitted by clients - for a given cron expression.
2. Schedule jobs using a fixed pool of worker.
3. Support concurrent execution of jobs with a maximum concurrency limit.
4. Allow job cancellation.
5. Return results/errors of each job

### Core Components

1. Job
- Attributes => JobId, Task, CronString, NextScheduledTime, Status : Cancelled-1, Active-0
- Methods => CancelJob(), UpdateNextScheduledTime()

2. ScheduledTask 
- Attributes => ScheduledTaskId, JobId, Status: Scheduled-1, Failed-2, Callenced-3, InProgress-4, Completed-5, TimeOfExecution, Retry, ErrorMessage
- JobId + TimeOfExecution -> unique key
- Methods => UpdateStatus()

3. Worker 
- Attributes => WorkerId, CurrentTaskAssigned, Status: Free-0, Working-1, TimeOfTaskAssigned
- Methods => ExecuteTask(), StopExecution()

4. JobManager
- Attributes => []Job
- Methods => 
 - CheckJobsNextSchedule()
>   Loop over Job List <br>
>   - Step1 : if NextScheduledTime <=  CurrentTime <br>
>   - Step2 : Put it in ScheduledTask List <br>
>   - Step3: Update nextScheduledTime <br>

5. Scheduler
- Attributes =>	[]Worker, []ScheduledTask, FreeWorkerIds, BusyWorkersIds, WorkerTaskMapping: map{taskId}workerId
- Methods =>
 - GetScheduledTasks(): Loop over scheduled tasks where status is scheduled 
 - AssignTask() 
    > - tasks = GetScheduledTasks()<br>
	> - check free workers list<br>
	> - assign given tasks and put in WorkerTaskMapping<br>
	> - Change status of worker : working and task status: In progress and call ExecuteTask() for each worker<br>
 - CheckTaskCompletion() 
    > Check WorkerTaskMapping for each running task, if task assignment is more than 10 mins call StopExecution()<br>
 - TaskCompleted(workerId, taskId uint64) {
	> worker status : free, task status : completed <br>
	> add worker to free workers list
 - CancelScheduledTask(jobId int)
	> update in scheduled tasks if status of jobId  = Scheduled, In Progress ->Cancelled <br>
	> Check taskId in Controller WorkerTaskMapping and call stop execution
