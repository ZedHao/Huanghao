TESTING_PREFIX=offline  go test -cover -race -v -count=1 ./service/...  -run="TestCheckInSuite/TestBatchGetActivityThreadCount" 
 
 2. mock的常见报错
 
 2.media_test.go:163: wrong type of argument 0 to Return for *dbmock.MockMediaPenalty.FindMediaPenaltyCount: int is not assignable to int64 
 
 	s.mockMediaPenalty.EXPECT().FindMediaPenaltyCount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(10), errors.New("test")).AnyTimes()
