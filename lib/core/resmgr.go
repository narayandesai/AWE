package core

import ()

type ClientMgr interface {
	RegisterNewClient(FormFiles) (*Client, error)
	ClientHeartBeat(string) (HBmsg, error)
	GetClient(string) (*Client, error)
	GetAllClients() []*Client
	DeleteClient(string)
	SuspendClient(string) (err error)
	ResumeClient(string) (err error)
	ResumeSuspendedClients() (count int)
	SuspendAllClients() (count int)
	ClientChecker()
	UpdateSubClients(id string, count int)
}

type WorkMgr interface {
	GetWorkById(string) (*Workunit, error)
	ShowWorkunits(string) []*Workunit
	CheckoutWorkunits(string, string, int) ([]*Workunit, error)
	NotifyWorkStatus(Notice)
	EnqueueWorkunit(*Workunit) error
	FetchDataToken(string, string) (string, error)
}

type JobMgr interface {
	JobRegister() (string, error)
	EnqueueTasksByJobId(string, []*Task) error
	GetActiveJobs() map[string]*JobPerf
	GetSuspendJobs() map[string]bool
	SuspendJob(string, string) error
	ResumeSuspendedJob(string) error
	ResumeSuspendedJobs() int
	ResubmitJob(string) error
	DeleteJob(string) error
	DeleteSuspendedJobs() int
	DeleteZombieJobs() int
	InitMaxJid() error
	RecoverJobs() error
	FinalizeWorkPerf(string, string) error
	RecomputeJob(string, string) error
	UpdateGroup(string, string) error
}

type ClientWorkMgr interface {
	ClientMgr
	WorkMgr
}

type ResourceMgr interface {
	ClientWorkMgr
	JobMgr
	Handle()
	ShowStatus() QueueStatus
	Timer()
}
