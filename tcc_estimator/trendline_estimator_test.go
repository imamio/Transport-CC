package tccEstimator

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type TrendlineIn struct {
	recv_time float64
	send_time float64
	arr_time  int64
}

type TrendlineOut struct {
	param1 float64
	param2 float64
	param3 float64
	param4 float64
	param5 float64
}

func TestTrendlineFilter(t *testing.T) {
	trendline_estimator := NewTrendlineEstimator()
	input := []TrendlineIn{
		TrendlineIn{recv_time: 19, send_time: 17.9977, arr_time: 1228983306},
		TrendlineIn{recv_time: 11, send_time: 11.0016, arr_time: 1228983313},
		TrendlineIn{recv_time: 7, send_time: 6.00052, arr_time: 1228983323},
		TrendlineIn{recv_time: 10, send_time: 9.99832, arr_time: 1228983338},
		TrendlineIn{recv_time: 18, send_time: 20.0005, arr_time: 1228983363},
		TrendlineIn{recv_time: 22, send_time: 21.9994, arr_time: 1228983383},
		TrendlineIn{recv_time: 20, send_time: 20.0005, arr_time: 1228983385},
		TrendlineIn{recv_time: 2, send_time: 0.999451, arr_time: 1228983395},
		TrendlineIn{recv_time: 10, send_time: 8.99887, arr_time: 1228983405},
		TrendlineIn{recv_time: 11, send_time: 11.0016, arr_time: 1228983476},
		TrendlineIn{recv_time: 72, send_time: 16.9983, arr_time: 1228983479},
		TrendlineIn{recv_time: 3, send_time: 53.0014, arr_time: 1228983482},
		TrendlineIn{recv_time: 5, send_time: 11.0016, arr_time: 1228983487},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228983512},
		TrendlineIn{recv_time: 25, send_time: 20.9999, arr_time: 1228983523},
		TrendlineIn{recv_time: 11, send_time: 14.9994, arr_time: 1228983565},
		TrendlineIn{recv_time: 71, send_time: 70.9991, arr_time: 1228983597},
		TrendlineIn{recv_time: 5, send_time: 4.00162, arr_time: 1228983619},
		TrendlineIn{recv_time: 22, send_time: 22.9988, arr_time: 1228983639},
		TrendlineIn{recv_time: 18, send_time: 13.0005, arr_time: 1228983652},
		TrendlineIn{recv_time: 33, send_time: 27.9999, arr_time: 1228983673},
		TrendlineIn{recv_time: 1, send_time: 2.99835, arr_time: 1228983692},
		TrendlineIn{recv_time: 23, send_time: 31.002, arr_time: 1228983712},
		TrendlineIn{recv_time: 16, send_time: 6.99997, arr_time: 1228983718},
		TrendlineIn{recv_time: 9, send_time: 16.9983, arr_time: 1228983734},
		TrendlineIn{recv_time: 50, send_time: 29.0031, arr_time: 1228983816},
		TrendlineIn{recv_time: 106, send_time: 29.9988, arr_time: 1228983913},
		TrendlineIn{recv_time: 36, send_time: 1.9989, arr_time: 1228983938},
		TrendlineIn{recv_time: 92, send_time: 34.0004, arr_time: 1228984045},
		TrendlineIn{recv_time: 72, send_time: 4.00162, arr_time: 1228984108},
		TrendlineIn{recv_time: 31, send_time: 6.99997, arr_time: 1228984139},
		TrendlineIn{recv_time: 111, send_time: 30.9982, arr_time: 1228984262},
		TrendlineIn{recv_time: 130, send_time: 48.0003, arr_time: 1228984359},
		TrendlineIn{recv_time: 43, send_time: 27.9999, arr_time: 1228984393},
		TrendlineIn{recv_time: 54, send_time: 9.99832, arr_time: 1228984481},
		TrendlineIn{recv_time: 36, send_time: 22.0032, arr_time: 1228984513},
		TrendlineIn{recv_time: 96, send_time: 29.9988, arr_time: 1228984624},
		TrendlineIn{recv_time: 109, send_time: 34.0004, arr_time: 1228984731},
		TrendlineIn{recv_time: 107, send_time: 28.9993, arr_time: 1228984825},
		TrendlineIn{recv_time: 91, send_time: 33.0009, arr_time: 1228984914},
		TrendlineIn{recv_time: 91, send_time: 10.9978, arr_time: 1228985007},
		TrendlineIn{recv_time: 62, send_time: 54.0009, arr_time: 1228985042},
		TrendlineIn{recv_time: 4, send_time: 0.999451, arr_time: 1228985071},
		TrendlineIn{recv_time: 61, send_time: 34.0004, arr_time: 1228985149},
		TrendlineIn{recv_time: 47, send_time: 76.9997, arr_time: 1228985166},
		TrendlineIn{recv_time: 47, send_time: 32.0015, arr_time: 1228985227},
		TrendlineIn{recv_time: 65, send_time: 63.9992, arr_time: 1228985263},
		TrendlineIn{recv_time: 31, send_time: 33.0009, arr_time: 1228985322},
		TrendlineIn{recv_time: 30, send_time: 63.9992, arr_time: 1228985367},
		TrendlineIn{recv_time: 77, send_time: 54.0009, arr_time: 1228985429},
		TrendlineIn{recv_time: 30, send_time: 41.9998, arr_time: 1228985459},
		TrendlineIn{recv_time: 29, send_time: 31.9977, arr_time: 1228985460},
		TrendlineIn{recv_time: 32, send_time: 45.002, arr_time: 1228985537},
		TrendlineIn{recv_time: 46, send_time: 32.0015, arr_time: 1228985538},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228985566},
		TrendlineIn{recv_time: 28, send_time: 28.9993, arr_time: 1228985597},
		TrendlineIn{recv_time: 31, send_time: 0.999451, arr_time: 1228985628},
		TrendlineIn{recv_time: 31, send_time: 29.9988, arr_time: 1228985647},
		TrendlineIn{recv_time: 20, send_time: 32.0015, arr_time: 1228985658},
		TrendlineIn{recv_time: 10, send_time: 32.0015, arr_time: 1228985688},
		TrendlineIn{recv_time: 31, send_time: 31.9977, arr_time: 1228985719},
		TrendlineIn{recv_time: 31, send_time: 32.0015, arr_time: 1228985751},
		TrendlineIn{recv_time: 31, send_time: 12.001, arr_time: 1228985782},
		TrendlineIn{recv_time: 31, send_time: 19.9966, arr_time: 1228985813},
		TrendlineIn{recv_time: 33, send_time: 44.0025, arr_time: 1228985862},
		TrendlineIn{recv_time: 49, send_time: 21.9994, arr_time: 1228985894},
		TrendlineIn{recv_time: 58, send_time: 54.0009, arr_time: 1228985924},
		TrendlineIn{recv_time: 3, send_time: 20.9999, arr_time: 1228985953},
		TrendlineIn{recv_time: 28, send_time: 41.9998, arr_time: 1228985955},
		TrendlineIn{recv_time: 3, send_time: 0.999451, arr_time: 1228985983},
		TrendlineIn{recv_time: 58, send_time: 61.0008, arr_time: 1228986045},
		TrendlineIn{recv_time: 31, send_time: 20.0005, arr_time: 1228986046},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228986076},
		TrendlineIn{recv_time: 31, send_time: 31.9977, arr_time: 1228986106},
		TrendlineIn{recv_time: 31, send_time: 23.0026, arr_time: 1228986137},
		TrendlineIn{recv_time: 29, send_time: 21.9994, arr_time: 1228986168},
		TrendlineIn{recv_time: 32, send_time: 33.0009, arr_time: 1228986206},
		TrendlineIn{recv_time: 37, send_time: 32.9971, arr_time: 1228986230},
		TrendlineIn{recv_time: 25, send_time: 34.9998, arr_time: 1228986262},
		TrendlineIn{recv_time: 31, send_time: 44.0025, arr_time: 1228986292},
		TrendlineIn{recv_time: 30, send_time: 13.9999, arr_time: 1228986323},
		TrendlineIn{recv_time: 31, send_time: 21.9994, arr_time: 1228986352},
		TrendlineIn{recv_time: 30, send_time: 30.9982, arr_time: 1228986383},
		TrendlineIn{recv_time: 30, send_time: 23.0026, arr_time: 1228986419},
		TrendlineIn{recv_time: 36, send_time: 31.9977, arr_time: 1228986445},
		TrendlineIn{recv_time: 27, send_time: 44.0025, arr_time: 1228986477},
		TrendlineIn{recv_time: 32, send_time: 29.9988, arr_time: 1228986505},
		TrendlineIn{recv_time: 27, send_time: 31.9977, arr_time: 1228986538},
		TrendlineIn{recv_time: 33, send_time: 32.0015, arr_time: 1228986568},
		TrendlineIn{recv_time: 30, send_time: 24.0021, arr_time: 1228986599},
		TrendlineIn{recv_time: 31, send_time: 22.9988, arr_time: 1228986629},
		TrendlineIn{recv_time: 32, send_time: 44.9982, arr_time: 1228986660},
		TrendlineIn{recv_time: 29, send_time: 33.0009, arr_time: 1228986661},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228986704},
		TrendlineIn{recv_time: 44, send_time: 32.0015, arr_time: 1228986735},
		TrendlineIn{recv_time: 31, send_time: 31.9977, arr_time: 1228986769},
		TrendlineIn{recv_time: 34, send_time: 45.002, arr_time: 1228986801},
		TrendlineIn{recv_time: 31, send_time: 33.0009, arr_time: 1228986829},
		TrendlineIn{recv_time: 29, send_time: 30.9982, arr_time: 1228986861},
		TrendlineIn{recv_time: 31, send_time: 21.9994, arr_time: 1228986892},
		TrendlineIn{recv_time: 31, send_time: 11.0016, arr_time: 1228986925},
		TrendlineIn{recv_time: 34, send_time: 67.0013, arr_time: 1228986954},
		TrendlineIn{recv_time: 29, send_time: 20.9999, arr_time: 1228986989},
		TrendlineIn{recv_time: 34, send_time: 21.9994, arr_time: 1228987013},
		TrendlineIn{recv_time: 25, send_time: 35.9993, arr_time: 1228987058},
		TrendlineIn{recv_time: 45, send_time: 30.9982, arr_time: 1228987089},
		TrendlineIn{recv_time: 31, send_time: 43.0031, arr_time: 1228987122},
		TrendlineIn{recv_time: 32, send_time: 30.9982, arr_time: 1228987153},
		TrendlineIn{recv_time: 31, send_time: 22.9988, arr_time: 1228987184},
		TrendlineIn{recv_time: 32, send_time: 33.0009, arr_time: 1228987216},
		TrendlineIn{recv_time: 31, send_time: 20.9999, arr_time: 1228987245},
		TrendlineIn{recv_time: 30, send_time: 56.9992, arr_time: 1228987276},
		TrendlineIn{recv_time: 31, send_time: 10.0021, arr_time: 1228987326},
		TrendlineIn{recv_time: 49, send_time: 21.9994, arr_time: 1228987337},
		TrendlineIn{recv_time: 12, send_time: 55.0003, arr_time: 1228987386},
		TrendlineIn{recv_time: 49, send_time: 46.0014, arr_time: 1228987440},
		TrendlineIn{recv_time: 54, send_time: 23.9983, arr_time: 1228987463},
		TrendlineIn{recv_time: 23, send_time: 33.0009, arr_time: 1228987495},
		TrendlineIn{recv_time: 32, send_time: 42.9993, arr_time: 1228987524},
		TrendlineIn{recv_time: 34, send_time: 44.9982, arr_time: 1228987556},
		TrendlineIn{recv_time: 27, send_time: 32.0015, arr_time: 1228987585},
		TrendlineIn{recv_time: 29, send_time: 11.0016, arr_time: 1228987619},
		TrendlineIn{recv_time: 33, send_time: 35.9993, arr_time: 1228987620},
		TrendlineIn{recv_time: 30, send_time: 42.9993, arr_time: 1228987679},
		TrendlineIn{recv_time: 31, send_time: 20.9999, arr_time: 1228987710},
		TrendlineIn{recv_time: 30, send_time: 36.9987, arr_time: 1228987741},
		TrendlineIn{recv_time: 31, send_time: 24.0021, arr_time: 1228987769},
		TrendlineIn{recv_time: 29, send_time: 31.9977, arr_time: 1228987802},
		TrendlineIn{recv_time: 33, send_time: 33.0009, arr_time: 1228987911},
		TrendlineIn{recv_time: 109, send_time: 122.002, arr_time: 1228987927},
		TrendlineIn{recv_time: 15, send_time: 42.9993, arr_time: 1228987954},
		TrendlineIn{recv_time: 28, send_time: 10.9978, arr_time: 1228987984},
		TrendlineIn{recv_time: 30, send_time: 20.9999, arr_time: 1228988015},
		TrendlineIn{recv_time: 30, send_time: 11.0016, arr_time: 1228988050},
		TrendlineIn{recv_time: 35, send_time: 54.0009, arr_time: 1228988076},
		TrendlineIn{recv_time: 27, send_time: 30.9982, arr_time: 1228988109},
		TrendlineIn{recv_time: 32, send_time: 32.0015, arr_time: 1228988139},
		TrendlineIn{recv_time: 31, send_time: 21.9994, arr_time: 1228988172},
		TrendlineIn{recv_time: 32, send_time: 32.0015, arr_time: 1228988201},
		TrendlineIn{recv_time: 30, send_time: 85.9985, arr_time: 1228988232},
		TrendlineIn{recv_time: 30, send_time: 45.002, arr_time: 1228988260},
		TrendlineIn{recv_time: 29, send_time: 10.9978, arr_time: 1228988292},
		TrendlineIn{recv_time: 32, send_time: 40.0009, arr_time: 1228988321},
		TrendlineIn{recv_time: 29, send_time: 43.9987, arr_time: 1228988352},
		TrendlineIn{recv_time: 30, send_time: 23.0026, arr_time: 1228988382},
		TrendlineIn{recv_time: 30, send_time: 60.997, arr_time: 1228988413},
		TrendlineIn{recv_time: 31, send_time: 22.0032, arr_time: 1228988443},
		TrendlineIn{recv_time: 32, send_time: 10.9978, arr_time: 1228988472},
		TrendlineIn{recv_time: 58, send_time: 53.0014, arr_time: 1228988533},
		TrendlineIn{recv_time: 31, send_time: 10.9978, arr_time: 1228988569},
		TrendlineIn{recv_time: 35, send_time: 33.0009, arr_time: 1228988594},
		TrendlineIn{recv_time: 26, send_time: 21.9994, arr_time: 1228988625},
		TrendlineIn{recv_time: 30, send_time: 30.0026, arr_time: 1228988655},
		TrendlineIn{recv_time: 31, send_time: 10.9978, arr_time: 1228988686},
		TrendlineIn{recv_time: 30, send_time: 20.9999, arr_time: 1228988687},
		TrendlineIn{recv_time: 30, send_time: 13.9999, arr_time: 1228988717},
		TrendlineIn{recv_time: 32, send_time: 130.001, arr_time: 1228988775},
		TrendlineIn{recv_time: 29, send_time: 10.9978, arr_time: 1228988807},
		TrendlineIn{recv_time: 30, send_time: 10.0021, arr_time: 1228988808},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228988837},
		TrendlineIn{recv_time: 29, send_time: 20.0005, arr_time: 1228988869},
		TrendlineIn{recv_time: 32, send_time: 13.0005, arr_time: 1228988898},
		TrendlineIn{recv_time: 30, send_time: 20.0005, arr_time: 1228988927},
		TrendlineIn{recv_time: 29, send_time: 10.9978, arr_time: 1228988961},
		TrendlineIn{recv_time: 34, send_time: 53.0014, arr_time: 1228988992},
		TrendlineIn{recv_time: 30, send_time: 13.0005, arr_time: 1228989021},
		TrendlineIn{recv_time: 30, send_time: 23.9983, arr_time: 1228989051},
		TrendlineIn{recv_time: 30, send_time: 21.9994, arr_time: 1228989083},
		TrendlineIn{recv_time: 31, send_time: 41.0004, arr_time: 1228989113},
		TrendlineIn{recv_time: 31, send_time: 34.9998, arr_time: 1228989143},
		TrendlineIn{recv_time: 30, send_time: 23.0026, arr_time: 1228989175},
		TrendlineIn{recv_time: 34, send_time: 32.9971, arr_time: 1228989205},
		TrendlineIn{recv_time: 28, send_time: 23.0026, arr_time: 1228989236},
		TrendlineIn{recv_time: 32, send_time: 56.9992, arr_time: 1228989269},
		TrendlineIn{recv_time: 31, send_time: 20.9999, arr_time: 1228989302},
		TrendlineIn{recv_time: 33, send_time: 69.0002, arr_time: 1228989330},
		TrendlineIn{recv_time: 29, send_time: 20.9999, arr_time: 1228989360},
		TrendlineIn{recv_time: 30, send_time: 43.9987, arr_time: 1228989391},
		TrendlineIn{recv_time: 30, send_time: 32.0015, arr_time: 1228989420},
		TrendlineIn{recv_time: 29, send_time: 51.9981, arr_time: 1228989453},
		TrendlineIn{recv_time: 33, send_time: 15.0032, arr_time: 1228989484},
		TrendlineIn{recv_time: 62, send_time: 85.9985, arr_time: 1228989547},
		TrendlineIn{recv_time: 32, send_time: 13.0005, arr_time: 1228989577},
		TrendlineIn{recv_time: 61, send_time: 29.9988, arr_time: 1228989609},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228989638},
		TrendlineIn{recv_time: 30, send_time: 20.0005, arr_time: 1228989674},
		TrendlineIn{recv_time: 53, send_time: 47.0009, arr_time: 1228989701},
		TrendlineIn{recv_time: 10, send_time: 10.9978, arr_time: 1228989730},
		TrendlineIn{recv_time: 29, send_time: 20.0005, arr_time: 1228989761},
		TrendlineIn{recv_time: 30, send_time: 32.0015, arr_time: 1228989790},
		TrendlineIn{recv_time: 29, send_time: 11.0016, arr_time: 1228989820},
		TrendlineIn{recv_time: 31, send_time: 26.9966, arr_time: 1228989865},
		TrendlineIn{recv_time: 94, send_time: 147.003, arr_time: 1228989916},
		TrendlineIn{recv_time: 1, send_time: 0.999451, arr_time: 1228989949},
		TrendlineIn{recv_time: 33, send_time: 56.9992, arr_time: 1228989977},
		TrendlineIn{recv_time: 29, send_time: 64.9986, arr_time: 1228990011},
		TrendlineIn{recv_time: 35, send_time: 35.9993, arr_time: 1228990041},
		TrendlineIn{recv_time: 28, send_time: 1.00327, arr_time: 1228990077},
		TrendlineIn{recv_time: 37, send_time: 30.9982, arr_time: 1228990104},
		TrendlineIn{recv_time: 27, send_time: 29.9988, arr_time: 1228990134},
	}
	output := []TrendlineOut{
		TrendlineOut{param1: 0, param2: 0.100226, param3: 1.00226, param4: 0, param5: 12.5},
		TrendlineOut{param1: 0, param2: 0.19027, param3: 1.00067, param4: 0, param5: 12.5},
		TrendlineOut{param1: 0, param2: 0.371259, param3: 2.00015, param4: 0, param5: 7.625},
		TrendlineOut{param1: 0, param2: 0.534316, param3: 2.00183, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 0.481022, param3: 0.00137329, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 0.433121, param3: 0.00201416, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 0.389964, param3: 0.0015564, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 0.451179, param3: 1.00211, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 0.606384, param3: 2.00323, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 0.745911, param3: 2.00165, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 6.37166, param3: 57.0034, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 6.43468, param3: 7.00195, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 5.89125, param3: 1.00037, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 5.40222, param3: 1.00092, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 5.3621, param3: 5.00101, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 4.92605, param3: 1.00162, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 4.53369, param3: 1.00247, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 4.28041, param3: 2.00085, param4: 0, param5: 6},
		TrendlineOut{param1: 0, param2: 3.95257, param3: 1.00204, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0163985, param2: 4.15747, param3: 6.00156, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0152516, param2: 4.84189, param3: 11.0017, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0143336, param2: 5.25803, param3: 9.00333, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0129959, param2: 4.83236, param3: 1.00128, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0120511, param2: 5.34925, param3: 10.0013, param4: 0, param5: 6},
		TrendlineOut{param1: 0.010515, param2: 5.01463, param3: 2.00302, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00949503, param2: 6.81316, param3: 22.9999, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0156087, param2: 16.032, param3: 99.0011, param4: 0, param5: 6},
		TrendlineOut{param1: 0.027842, param2: 27.729, param3: 133.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0476071, param2: 44.0563, param3: 191.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.070606, param2: 65.5507, param3: 259, param4: 0, param5: 7.3553},
		TrendlineOut{param1: 0.0969335, param2: 87.2956, param3: 283, param4: 0, param5: 8.6133},
		TrendlineOut{param1: 0.12512, param2: 114.866, param3: 363.002, param4: 2, param5: 15.0531},
		TrendlineOut{param1: 0.154203, param2: 147.88, param3: 445.002, param4: 2, param5: 19.5272},
		TrendlineOut{param1: 0.182864, param2: 179.092, param3: 460.002, param4: 2, param5: 21.1074},
		TrendlineOut{param1: 0.210457, param2: 211.583, param3: 504.003, param4: 2, param5: 27.5052},
		TrendlineOut{param1: 0.236344, param2: 242.225, param3: 518, param4: 2, param5: 29.3227},
		TrendlineOut{param1: 0.258551, param2: 276.403, param3: 584.001, param4: 2, param5: 37.1029},
		TrendlineOut{param1: 0.278441, param2: 314.662, param3: 659.001, param4: 2, param5: 41.6445},
		TrendlineOut{param1: 0.297641, param2: 356.896, param3: 737.002, param4: 2, param5: 45.5597},
		TrendlineOut{param1: 0.317054, param2: 400.707, param3: 795.001, param4: 2, param5: 49.562},
		TrendlineOut{param1: 0.336591, param2: 448.136, param3: 875.003, param4: 2, param5: 54.1245},
		TrendlineOut{param1: 0.358681, param2: 491.623, param3: 883.002, param4: 2, param5: 55.9923},
		TrendlineOut{param1: 0.382688, param2: 531.061, param3: 886.003, param4: 2, param5: 58.4724},
		TrendlineOut{param1: 0.407889, param2: 569.255, param3: 913.002, param4: 2, param5: 67.5087},
		TrendlineOut{param1: 0.436917, param2: 600.63, param3: 883.003, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.461127, param2: 630.367, param3: 898.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.480519, param2: 657.23, param3: 899.002, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.501525, param2: 681.208, param3: 897.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.514985, param2: 699.387, param3: 863.002, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.525269, param2: 718.048, param3: 886.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.540126, param2: 733.644, param3: 874.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.549412, param2: 747.38, param3: 871.004, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.548209, param2: 758.442, param3: 858.002, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.554851, param2: 769.798, param3: 872, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.557154, param2: 780.018, param3: 872.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.56559, param2: 789.116, param3: 871.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.563332, param2: 800.305, param3: 901.002, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.551174, param2: 810.475, param3: 902.003, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.532713, param2: 818.427, param3: 890.002, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.504999, param2: 823.385, param3: 868, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.470237, param2: 827.746, param3: 867.002, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.437879, param2: 831.572, param3: 866.001, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.409815, param2: 836.915, param3: 885, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.379331, param2: 842.824, param3: 896.003, param4: 2, param5: 69.1558},
		TrendlineOut{param1: 0.350563, param2: 847.041, param3: 885.001, param4: 2, param5: 75.5415},
		TrendlineOut{param1: 0.322241, param2: 853.537, param3: 912.001, param4: 2, param5: 76.0416},
		TrendlineOut{param1: 0.298884, param2: 859.784, param3: 916, param4: 0, param5: 70.9995},
		TrendlineOut{param1: 0.276093, param2: 863.605, param3: 898.001, param4: 0, param5: 65.6416},
		TrendlineOut{param1: 0.257155, param2: 865.645, param3: 884.001, param4: 0, param5: 65.3355},
		TrendlineOut{param1: 0.237242, param2: 867.68, param3: 886.001, param4: 0, param5: 56.1656},
		TrendlineOut{param1: 0.213974, param2: 869.212, param3: 883, param4: 0, param5: 44.5306},
		TrendlineOut{param1: 0.198838, param2: 871.691, param3: 894, param4: 0, param5: 44.5584},
		TrendlineOut{param1: 0.181099, param2: 873.922, param3: 894.001, param4: 0, param5: 43.2778},
		TrendlineOut{param1: 0.165814, param2: 875.83, param3: 893.003, param4: 0, param5: 39.2032},
		TrendlineOut{param1: 0.151651, param2: 878.347, param3: 901, param4: 0, param5: 35.8095},
		TrendlineOut{param1: 0.138589, param2: 881.313, param3: 908.001, param4: 0, param5: 32.7287},
		TrendlineOut{param1: 0.127384, param2: 883.881, param3: 907, param4: 0, param5: 29.5327},
		TrendlineOut{param1: 0.119579, param2: 886.593, param3: 911.003, param4: 0, param5: 28.7524},
		TrendlineOut{param1: 0.113545, param2: 888.034, param3: 901.003, param4: 0, param5: 26.8784},
		TrendlineOut{param1: 0.106535, param2: 888.031, param3: 888, param4: 0, param5: 25.3457},
		TrendlineOut{param1: 0.0991241, param2: 889.628, param3: 904, param4: 0, param5: 23.4646},
		TrendlineOut{param1: 0.0916179, param2: 891.965, param3: 913.001, param4: 0, param5: 21.7949},
		TrendlineOut{param1: 0.0847187, param2: 893.969, param3: 912.003, param4: 0, param5: 20.0268},
		TrendlineOut{param1: 0.0789951, param2: 896.472, param3: 919, param4: 0, param5: 18.5274},
		TrendlineOut{param1: 0.0738453, param2: 899.125, param3: 923.003, param4: 0, param5: 17.7116},
		TrendlineOut{param1: 0.0697393, param2: 899.813, param3: 906, param4: 0, param5: 16.4958},
		TrendlineOut{param1: 0.0670401, param2: 900.632, param3: 908.001, param4: 0, param5: 16.0522},
		TrendlineOut{param1: 0.0643246, param2: 900.869, param3: 903.004, param4: 0, param5: 15.2616},
		TrendlineOut{param1: 0.06181, param2: 901.182, param3: 904.002, param4: 0, param5: 14.7618},
		TrendlineOut{param1: 0.0590592, param2: 902.064, param3: 910, param4: 0, param5: 14.0514},
		TrendlineOut{param1: 0.0554697, param2: 903.658, param3: 918.001, param4: 0, param5: 13.1872},
		TrendlineOut{param1: 0.052151, param2: 903.792, param3: 905.003, param4: 0, param5: 12.376},
		TrendlineOut{param1: 0.0490997, param2: 903.513, param3: 901.002, param4: 0, param5: 12.3529},
		TrendlineOut{param1: 0.0449541, param2: 903.262, param3: 901.003, param4: 0, param5: 9.73021},
		TrendlineOut{param1: 0.0410703, param2: 904.236, param3: 913.001, param4: 0, param5: 9.76437},
		TrendlineOut{param1: 0.0376413, param2: 905.013, param3: 912.003, param4: 0, param5: 8.79579},
		TrendlineOut{param1: 0.0340058, param2: 904.612, param3: 901.001, param4: 0, param5: 8.00406},
		TrendlineOut{param1: 0.0306015, param2: 904.051, param3: 899.001, param4: 0, param5: 7.28366},
		TrendlineOut{param1: 0.0266408, param2: 903.346, param3: 897.002, param4: 0, param5: 6.17311},
		TrendlineOut{param1: 0.0219849, param2: 903.611, param3: 906.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.018446, param2: 905.85, param3: 926.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0147289, param2: 904.565, param3: 893, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0112367, param2: 904.209, param3: 901, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00907433, param2: 905.088, param3: 913.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00766003, param2: 904.779, param3: 902.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00690952, param2: 905.902, param3: 916.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00620046, param2: 905.712, param3: 904, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00535329, param2: 905.641, param3: 905.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00480592, param2: 906.377, param3: 913.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00471048, param2: 906.939, param3: 912.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00581399, param2: 908.446, param3: 922.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00613415, param2: 907.101, param3: 895.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00662585, param2: 907.991, param3: 916.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00831557, param2: 911.492, param3: 943.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00957934, param2: 910.343, param3: 900.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0104589, param2: 909.609, param3: 903, param4: 0, param5: 6},
		TrendlineOut{param1: 0.012006, param2: 911.948, param3: 933.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.013443, param2: 913.053, param3: 923.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0141067, param2: 912.948, param3: 912.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0139742, param2: 911.754, param3: 901.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0138852, param2: 910.179, param3: 896.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0131968, param2: 910.561, param3: 914, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0123254, param2: 910.605, param3: 911.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0108511, param2: 909.344, param3: 898.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00892395, param2: 909.21, param3: 908.002, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00698744, param2: 908.389, param3: 901.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00483834, param2: 908.351, param3: 908.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0023483, param2: 908.016, param3: 905.003, param4: 0, param5: 6},
		TrendlineOut{param1: -0.000268556, param2: 907.714, param3: 905.002, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00314305, param2: 906.143, param3: 892.001, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00701593, param2: 901.929, param3: 864.001, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0120008, param2: 899.836, param3: 881.003, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0163627, param2: 898.853, param3: 890.004, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0185765, param2: 899.868, param3: 909.002, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0210811, param2: 898.881, param3: 890.001, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0236117, param2: 897.593, param3: 886.003, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0252649, param2: 896.434, param3: 886.001, param4: 1, param5: 6.01659},
		TrendlineOut{param1: -0.0259072, param2: 896.291, param3: 895.002, param4: 1, param5: 6.07434},
		TrendlineOut{param1: -0.0260761, param2: 896.162, param3: 895.001, param4: 1, param5: 6.12075},
		TrendlineOut{param1: -0.0282851, param2: 890.446, param3: 839.002, param4: 1, param5: 6.30082},
		TrendlineOut{param1: -0.0329615, param2: 883.801, param3: 824, param4: 1, param5: 6.693},
		TrendlineOut{param1: -0.0381135, param2: 879.621, param3: 842.002, param4: 1, param5: 7.37626},
		TrendlineOut{param1: -0.0447975, param2: 875.059, param3: 834.001, param4: 1, param5: 8.22781},
		TrendlineOut{param1: -0.0526503, param2: 869.454, param3: 819.003, param4: 1, param5: 9.41671},
		TrendlineOut{param1: -0.0612144, param2: 865.108, param3: 826, param4: 1, param5: 10.7934},
		TrendlineOut{param1: -0.071988, param2: 858.098, param3: 795.003, param4: 1, param5: 12.5421},
		TrendlineOut{param1: -0.084162, param2: 852.688, param3: 804, param4: 1, param5: 14.5405},
		TrendlineOut{param1: -0.0965575, param2: 849.919, param3: 825.002, param4: 1, param5: 16.7187},
		TrendlineOut{param1: -0.101286, param2: 847.928, param3: 830.001, param4: 1, param5: 20.7466},
		TrendlineOut{param1: -0.104097, param2: 848.135, param3: 850.003, param4: 1, param5: 22.0735},
		TrendlineOut{param1: -0.106548, param2: 848.522, param3: 852.002, param4: 1, param5: 22.8343},
		TrendlineOut{param1: -0.107249, param2: 849.27, param3: 856.003, param4: 1, param5: 23.6179},
		TrendlineOut{param1: -0.105931, param2: 849.943, param3: 856, param4: 1, param5: 24.0892},
		TrendlineOut{param1: -0.100646, param2: 852.549, param3: 876.002, param4: 1, param5: 24.107},
		TrendlineOut{param1: -0.0945587, param2: 855.794, param3: 885.002, param4: 0, param5: 24.0519},
		TrendlineOut{param1: -0.0849742, param2: 860.315, param3: 901.002, param4: 0, param5: 19.7719},
		TrendlineOut{param1: -0.0749358, param2: 854.584, param3: 803.001, param4: 0, param5: 15.729},
		TrendlineOut{param1: -0.0646428, param2: 851.226, param3: 821.004, param4: 0, param5: 15.461},
		TrendlineOut{param1: -0.053941, param2: 850.203, param3: 841.001, param4: 0, param5: 15.3629},
		TrendlineOut{param1: -0.0439618, param2: 849.283, param3: 841.002, param4: 0, param5: 9.92046},
		TrendlineOut{param1: -0.0349296, param2: 849.355, param3: 850.001, param4: 0, param5: 8.00185},
		TrendlineOut{param1: -0.0251864, param2: 851.319, param3: 869.001, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0147839, param2: 854.088, param3: 879.001, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00385907, param2: 858.379, param3: 897.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0066398, param2: 860.341, param3: 878.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0154875, param2: 863.807, param3: 895.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0231114, param2: 867.527, param3: 901.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0307804, param2: 871.674, param3: 909.003, param4: 0, param5: 6.38622},
		TrendlineOut{param1: 0.0369092, param2: 874.407, param3: 899.003, param4: 2, param5: 7.03141},
		TrendlineOut{param1: 0.0424897, param2: 876.467, param3: 895.003, param4: 2, param5: 7.85777},
		TrendlineOut{param1: 0.0479962, param2: 879.02, param3: 902, param4: 2, param5: 8.87708},
		TrendlineOut{param1: 0.0532772, param2: 881.419, param3: 903.003, param4: 2, param5: 9.89745},
		TrendlineOut{param1: 0.0583303, param2: 884.077, param3: 908.001, param4: 2, param5: 11.0037},
		TrendlineOut{param1: 0.0625592, param2: 883.969, param3: 883.001, param4: 2, param5: 12.1551},
		TrendlineOut{param1: 0.0684208, param2: 884.872, param3: 893.002, param4: 2, param5: 13.3799},
		TrendlineOut{param1: 0.0740303, param2: 882.085, param3: 857.001, param4: 2, param5: 14.4486},
		TrendlineOut{param1: 0.0731728, param2: 880.377, param3: 865.001, param4: 2, param5: 15.2611},
		TrendlineOut{param1: 0.0680339, param2: 877.44, param3: 851.003, param4: 2, param5: 15.5489},
		TrendlineOut{param1: 0.0610135, param2: 874.596, param3: 849.001, param4: 0, param5: 14.5246},
		TrendlineOut{param1: 0.0500119, param2: 869.736, param3: 826.003, param4: 0, param5: 11.2791},
		TrendlineOut{param1: 0.036864, param2: 867.163, param3: 844, param4: 0, param5: 8.33911},
		TrendlineOut{param1: 0.0205377, param2: 862.447, param3: 820.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00484274, param2: 860.102, param3: 839.001, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00781608, param2: 861.092, param3: 870.002, param4: 0, param5: 6},
		TrendlineOut{param1: -0.01872, param2: 861.983, param3: 870.003, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0269613, param2: 863.785, param3: 880.002, param4: 1, param5: 6.14742},
		TrendlineOut{param1: -0.03241, param2: 866.007, param3: 886.001, param4: 1, param5: 6.53054},
		TrendlineOut{param1: -0.0350428, param2: 867.906, param3: 885.003, param4: 1, param5: 7.00479},
		TrendlineOut{param1: -0.0352192, param2: 870.516, param3: 894.003, param4: 1, param5: 7.39527},
		TrendlineOut{param1: -0.0335308, param2: 872.665, param3: 892.001, param4: 1, param5: 7.5598},
		TrendlineOut{param1: -0.0290026, param2: 876.398, param3: 910, param4: 0, param5: 6.85875},
		TrendlineOut{param1: -0.0212989, param2: 880.159, param3: 914.003, param4: 0, param5: 6},
		TrendlineOut{param1: -0.0128675, param2: 878.243, param3: 861, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00531083, param2: 876.519, param3: 861.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00102782, param2: 872.567, param3: 837.001, param4: 0, param5: 6},
		TrendlineOut{param1: 0.0036246, param2: 865.41, param3: 801.003, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00354804, param2: 858.87, param3: 800.004, param4: 0, param5: 6},
		TrendlineOut{param1: 0.00174794, param2: 855.683, param3: 827, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00120072, param2: 853.415, param3: 833.002, param4: 0, param5: 6},
		TrendlineOut{param1: -0.00653747, param2: 851.074, param3: 830.003, param4: 0, param5: 6},
	}
	for i, value := range input {
		trendline_estimator.UpdateTrendline(value.recv_time, value.send_time, 0, value.arr_time, 0)

		//fmt.Println("smoothed_delay_: ", trendline_estimator.smooth_delay_, " acc: ", trendline_estimator.accumulated_delay_, " status ", trendline_estimator.hypothesis_, " threshold ", trendline_estimator.threshold_)
		assert.Equal(t, math.Abs(trendline_estimator.record_trend-output[i].param1) < 0.01, true)
		assert.Equal(t, math.Abs(trendline_estimator.smooth_delay_-output[i].param2) < 0.01, true)
		assert.Equal(t, math.Abs(trendline_estimator.accumulated_delay_-output[i].param3) < 0.01, true)
		assert.Equal(t, math.Abs(float64(trendline_estimator.hypothesis_)-output[i].param4) < 0.01, true)
		assert.Equal(t, math.Abs(trendline_estimator.threshold_-output[i].param5) < 0.01, true)

	}
}
