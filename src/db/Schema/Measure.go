package Schema

var KeyText map[string]string

type MeasureSimple struct{
	Id                                           string
	Name                                         string
	Company                                      string
	Address                                      string
	Phone                                        string
}

type Measure struct {
	Id                                           string
	Name                                         string
	Company                                      string
	Address                                      string
	Phone                                        string
	PhoneType                                    string
	MachineStruct                                string
	DesignUseSpeed                               string
	WorkTripX                                    string
	WorkTripY                                    string
	WorkTripZ                                    string
	ShieldBracketLengthX                         string
	ShieldBracketLengthY                         string
	ShieldBracketLengthZ                         string
	ShieldBracketLengthYH                        string
	ShieldStruct                                 string
	ShieldStructText                             string
	WorkBenchWidth                               string
	WorkBenchLength                              string
	TrackSurfaceToBenchBottomHeight              string
	TrackSurfaceToBenchTwoHeight                 string
	TrackSurfaceBenchTopHeight                   string
	WorkBenchMiddleDistanceBefore                string
	WorkBenchMiddleDistanceAfter                 string
	WorkBenchProtrudingSaddleDistance            string
	XInTrackWidth                                string
	XOutTrackWidth                               string
	TrackOutToTowlineSlotOutDistance             string
	TowlineHightLineTrackSurfaceHeight           string
	TrackOutToTowlineSlotInDistance              string
	XLineTrackSurfaceToMotorTopHeight            string
	XMotorWidth                                  string
	XLineTrackWidth                              string
	XLineTrackHeight                             string
	XLineTrackOutToCastingDistance               string
	XLineTrackInToCastingDistance                string
	XLineTrackInToMotorSideDistance              string
	YValidFixedSurfaceWidth                      string
	YValidFixedSurfaceHeight                     string
	TrackSurfaceToSaddleFixedSurfaceBottomHeight string
	TrackSurfaceToSaddleFixedSurfaceTopHeight    string
	YInTrackWidth                                string
	YOutTrackWidth                               string
	InsertsSideProtrudingDistance                string
	InsertsStretchSaddleDistance                 string
	NozzleProtrudingSaddleDistance               string
	NozzleHightTrackSurfaceDistance              string
	YMotorPlace                                  string
	YLineTrackSurfaceToMotorTopHeight            string
	YMotorWidth                                  string
	YLineTrackWidth                              string
	YLineTrackHeight                             string
	YLineTrackOutToCastingDistance               string
	YLineTrackInToCastingDistance                string
	YLineTrackInToMotorSideDistance              string
	YHStruct                                     string
	FixedSurfaceWidth                            string
	FixedSurfaceTopToTrackHeight                 string
	BottomSideToInBorderDistance                 string
	BottomTrackWidth                             string
	ColumnOpenSize                               string
	TrackToOpenTopHeight                         string
	OpenTopLength                                string
	YHLineTrackWidth                             string
	YHLineTrackHeight                            string
	YHLineTrackOutToCastingDistance              string
	YHLineTrackInToCastingDistance               string
	YHLineTrackInToMotorSideDistance             string
	YHLineTrackSurfaceToBottomHeight             string
	ZValidFixedSurfaceWidth                      string
	ZValidFixedSurfaceHeight                     string
	TrackToFixedTopHeight                        string
	ZInTrackWidth                                string
	ZOutTrackWidth                               string
	ZLineTrackWidth                              string
	ZLineTrackHeight                             string
	ZLineTrackOutToCastingDistance               string
	ZLineTrackInToCastingDistance                string
	ZLineTrackInToBearingSideDistance            string
	ZLineTrackToBearingTopHeight                 string
}

func init() {
	KeyText = make(map[string]string)
	KeyText["basicInfo"] = "基本信息"
	KeyText["name"] = "机型"
	KeyText["company"] = "机型所属公司"
	KeyText["address"] = "公司地址"
	KeyText["phone"] = "联系人以及联系电话"
	KeyText["phoneType"] = "联系人类型"
	KeyText["basicParam"] = "基本参数"
	KeyText["machineStruct"] = "机床结构"
	KeyText["designUseSpeed"] = "设计使用速度"
	KeyText["workTripX"] = "工作行程X轴"
	KeyText["workTripY"] = "工作行程Y轴"
	KeyText["workTripZ"] = "工作行程Z轴"
	KeyText["shieldBracketLengthX"] = "护罩支架长度X轴"
	KeyText["shieldBracketLengthY"] = "护罩支架长度Y轴"
	KeyText["shieldBracketLengthZ"] = "护罩支架长度Z轴"
	KeyText["shieldBracketLengthYH"] = "护罩支架长度Y后"
	KeyText["shieldStruct"] = "护罩结构要求"
	KeyText["shieldStructText"] = "护罩结构要求内容"
	KeyText["XPathMeasure"] = "X向测量"
	KeyText["workBench"] = "工作台"
	KeyText["workBenchWidth"] = "工作台宽度"
	KeyText["workBenchLength"] = "工作台长度"
	KeyText["trackSurfaceToBenchBottomHeight"] = "轨道面到工作台底端高度"
	KeyText["trackSurfaceToBenchTwoHeight"] = "轨道面到工作台第二台阶高度"
	KeyText["trackSurfaceBenchTopHeight"] = "轨道面到工作台顶面高度"
	KeyText["workBenchMiddleDistance"] = "工作台中分距离"
	KeyText["workBenchMiddleDistanceBefore"] = "前"
	KeyText["workBenchMiddleDistanceAfter"] = "后"
	KeyText["workBenchProtrudingSaddleDistance"] = "工作台凸出滑鞍距离"
	KeyText["xTrackWidth"] = "轨道宽度"
	KeyText["xInTrackWidth"] = "内轨宽"
	KeyText["xOutTrackWidth"] = "外轨宽"
	KeyText["trackOutToTowlineSlotOutDistance"] = "轨道外侧到拖链槽外边的距离"
	KeyText["towlineHightLineTrackSurfaceHeight"] = "拖链槽定边高出线轨面的高度"
	KeyText["trackOutToTowlineSlotInDistance"] = "轨道外侧到拖链槽内边的距离"
	KeyText["xLineTrackSurfaceToMotorTopHeight"] = "线轨面到电机座顶边的高度"
	KeyText["xMotorWidth"] = "电机座宽度"
	KeyText["xLineTrackWidth"] = "线轨的宽度"
	KeyText["xLineTrackHeight"] = "线轨的高度"
	KeyText["xLineTrackOutToCastingDistance"] = "线轨外侧到铸件边的距离"
	KeyText["xLineTrackInToCastingDistance"] = "线轨内侧到铸件边的距离"
	KeyText["xLineTrackInToMotorSideDistance"] = "线轨内侧到电机座侧边的距离"
	KeyText["YPathMeasure"] = "Y向测量"
	KeyText["saddleFixedSurface"] = "滑鞍固定面"
	KeyText["yValidFixedSurfaceWidth"] = "有效固定面的宽度"
	KeyText["yValidFixedSurfaceHeight"] = "有效固定面的高度"
	KeyText["trackSurfaceToSaddleFixedSurfaceBottomHeight"] = "轨道面到滑鞍固定面底端的高度"
	KeyText["trackSurfaceToSaddleFixedSurfaceTopHeight"] = "轨道面到滑鞍固定面顶端的高度"
	KeyText["yInTrackWidth"] = "内轨宽"
	KeyText["yOutTrackWidth"] = "外轨宽"
	KeyText["insertsSideProtrudingDistance"] = "镶条侧向凸出距离"
	KeyText["insertsStretchSaddleDistance"] = "镶条伸出滑鞍距离"
	KeyText["nozzleProtrudingSaddleDistance"] = "油嘴凸出滑鞍距离"
	KeyText["nozzleHightTrackSurfaceDistance"] = "油嘴高出轨道面的距离"
	KeyText["yMotorPlace"] = "电机放置"
	KeyText["yLineTrackSurfaceToMotorTopHeight"] = "线轨面到电机座顶边的高度"
	KeyText["yMotorWidth"] = "电机座宽度"
	KeyText["yLineTrackWidth"] = "线轨的宽度"
	KeyText["yLineTrackHeight"] = "线轨的高度"
	KeyText["yLineTrackOutToCastingDistance"] = "线轨外侧到铸件边的距离"
	KeyText["yLineTrackInToCastingDistance"] = "线轨内侧到铸件边的距离"
	KeyText["yLineTrackInToMotorSideDistance"] = "线轨内侧到电机座侧边的距离"
	KeyText["yHStruct"] = "Y后结构"
	KeyText["fixedSurfaceWidth"] = "固定面的宽度"
	KeyText["fixedSurfaceTopToTrackHeight"] = "固定面顶边到线轨面的高度"
	KeyText["bottomSideToInBorderDistance"] = "底座侧边到立柱内侧边缘的距离"
	KeyText["bottomTrackWidth"] = "底座线轨支座的宽度"
	KeyText["columnOpenSize"] = "立柱开裆尺寸"
	KeyText["trackToOpenTopHeight"] = "轨道面到开裆顶点高度"
	KeyText["openTopLength"] = "开裆顶边长度"
	KeyText["yHLineTrackWidth"] = "线轨的宽度"
	KeyText["yHLineTrackHeight"] = "线轨的高度"
	KeyText["yHLineTrackOutToCastingDistance"] = "线轨外侧到铸件边的距离"
	KeyText["yHLineTrackInToCastingDistance"] = "线轨内侧到铸件边的距离"
	KeyText["yHLineTrackInToMotorSideDistance"] = "线轨内侧到电机座侧边的距离"
	KeyText["yHLineTrackSurfaceToBottomHeight"] = "轨道面到立柱最低点的高度"
	KeyText["zValidFixedSurfaceWidth"] = "有效固定面的宽度"
	KeyText["zValidFixedSurfaceHeight"] = "有效固定面的高度"
	KeyText["trackToFixedTopHeight"] = "轨道面到固定面的顶端的高度"
	KeyText["zInTrackWidth"] = "内轨宽"
	KeyText["zOutTrackWidth"] = "外轨宽"
	KeyText["zLineTrackWidth"] = "线轨的宽度"
	KeyText["zLineTrackHeight"] = "线轨的高度"
	KeyText["zLineTrackOutToCastingDistance"] = "线轨外侧到铸件边的距离"
	KeyText["zLineTrackInToCastingDistance"] = "线轨内侧到铸件边的距离"
	KeyText["zLineTrackInToBearingSideDistance"] = "线轨内侧到轴承座侧边的距离"
	KeyText["zLineTrackToBearingTopHeight"] = "线轨面到轴承座顶边的高度"
}
