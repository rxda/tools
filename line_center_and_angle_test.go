package tools

import (
	"fmt"
	"github.com/go-courier/geography"
	"math"
	"reflect"
	"testing"
)

func TestGetCenterAndAngle(t *testing.T) {
	type args struct {
		lineString [2]geography.Point
	}
	tests := []struct {
		name  string
		args  args
		want  geography.Point
		want1 float64
	}{
		{
			name: "1象限45",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {1, 1}},
			},
			want:  geography.Point{0.5, 0.5},
			want1: 45,
		},
		{
			name: "1象限30",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {math.Sqrt(3), 1}},
			},
			want:  geography.Point{math.Sqrt(3) / 2, 0.5},
			want1: 30,
		},
		{
			name: "2象限",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {-1, 1}},
			},
			want:  geography.Point{-0.5, 0.5},
			want1: 135,
		},
		{
			name: "2象限30",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {-math.Sqrt(3), 1}},
			},
			want:  geography.Point{-math.Sqrt(3) / 2, 0.5},
			want1: 150,
		},
		{
			name: "3象限",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {-1, -1}},
			},
			want:  geography.Point{-0.5, -0.5},
			want1: 225,
		},
		{
			name: "3象限30",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {-math.Sqrt(3), -1}},
			},
			want:  geography.Point{-math.Sqrt(3) / 2, 0.5},
			want1: 210,
		},
		{
			name: "4象限",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {1, -1}},
			},
			want:  geography.Point{0.5, -0.5},
			want1: 315,
		},
		{
			name: "4象限30",
			args: args{
				lineString: [2]geography.Point{{0, 0}, {math.Sqrt(3), -1}},
			},
			want:  geography.Point{math.Sqrt(3) / 2, -0.5},
			want1: 330,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCenterAndAngle(tt.args.lineString)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCenterAndAngle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCenterAndAngle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

var data = [][2]geography.Point{

	{{106.3729325, 29.9325858}, {106.3746457, 29.9318720}},
	{{106.3762637, 29.9312057}, {106.3775010, 29.9305395}},
	{{106.3792618, 29.9294449}, {106.3808323, 29.9284456}},
	{{106.3823551, 29.9274938}, {106.3836876, 29.9264944}},
	{{106.3848298, 29.9249716}, {106.3856864, 29.9236391}},
	{{106.3863050, 29.9224017}, {106.3867333, 29.9207837}},
	{{106.3869713, 29.9190229}, {106.3866382, 29.9168814}},
	{{106.3860195, 29.9155013}, {106.3857816, 29.9135026}},
	{{106.3856864, 29.9119321}, {106.3856864, 29.9100285}},
	{{106.3859243, 29.9084581}, {106.3868285, 29.9069828}},
	{{106.3889938, 29.9048175}, {106.3897077, 29.9039847}},
	{{106.3903263, 29.9031757}, {106.3910402, 29.9023667}},
	{{106.3919444, 29.9013435}, {106.3926820, 29.9005821}},
	{{106.3937290, 29.8993686}, {106.3948235, 29.8987499}},
	{{106.3961322, 29.8982740}, {106.3969650, 29.8979409}},
	{{106.3981548, 29.8974412}, {106.3988686, 29.8971081}},
	{{106.3996776, 29.8964894}, {106.4004628, 29.8958945}},
	{{106.4012481, 29.8952283}, {106.4018667, 29.8946810}},
	{{106.4023426, 29.8938720}, {106.4025568, 29.8932058}},
	{{106.4025806, 29.8924919}, {106.4024616, 29.8914687}},
	{{106.4020571, 29.8906835}, {106.4015098, 29.8898269}},
	{{106.3996062, 29.8891131}, {106.3987734, 29.8883992}},
	{{106.4008198, 29.8874475}, {106.4002011, 29.8865909}},
	{{106.3977740, 29.8877568}, {106.3968223, 29.8870430}},
	{{106.3993921, 29.8856867}, {106.3987734, 29.8849728}},
	{{106.3975123, 29.8834024}, {106.3967509, 29.8826172}},
	{{106.3942372, 29.8807202}, {106.3934491, 29.8805603}},
	{{106.3937956, 29.8850337}, {106.3928438, 29.8842152}},
	{{106.3918730, 29.8823878}, {106.3914923, 29.8815121}},
	{{106.3916255, 29.8801987}, {106.3909783, 29.8792279}},
	{{106.3885455, 29.8818015}, {106.3881915, 29.8814474}},
	{{106.3871426, 29.8800326}, {106.3867619, 29.8794615}},
	{{106.3901883, 29.8782242}, {106.3896744, 29.8774628}},
	{{106.3862860, 29.8785288}, {106.3863050, 29.8778625}},
	{{106.3866477, 29.8766442}, {106.3872759, 29.8760922}},
	{{106.3891794, 29.8761303}, {106.3889510, 29.8754640}},
	{{106.3884180, 29.8750643}, {106.3888939, 29.8744171}},
	{{106.3892746, 29.8735605}, {106.3895601, 29.8728561}},
	{{106.3898076, 29.8720947}, {106.3901693, 29.8714285}},
	{{106.3905690, 29.8708764}, {106.3909688, 29.8703815}},
	{{106.3916541, 29.8699246}, {106.3920538, 29.8696581}},
	{{106.3929866, 29.8694107}, {106.3936147, 29.8692584}},
	{{106.3946427, 29.8692584}, {106.3954041, 29.8692584}},
	{{106.3964511, 29.8691632}, {106.3970412, 29.8690109}},
	{{106.3980881, 29.8687444}, {106.3987925, 29.8685541}},
	{{106.3995158, 29.8683066}, {106.4002011, 29.8681734}},
	{{106.4010196, 29.8680401}, {106.4018001, 29.8679069}},
	{{106.4028661, 29.8677165}, {106.4035704, 29.8675832}},
	{{106.4047316, 29.8671264}, {106.4057214, 29.8667837}},
	{{106.4070159, 29.8664792}, {106.4080057, 29.8659462}},
	{{106.4086910, 29.8653751}, {106.4094144, 29.8649182}},
	{{106.4102519, 29.8644233}, {106.4110895, 29.8637380}},
	{{106.4116225, 29.8633192}, {106.4122697, 29.8626720}},
	{{106.4127091, 29.8618963}, {106.4132802, 29.8611349}},
	{{106.4137561, 29.8603021}, {106.4141844, 29.8596834}},
	{{106.4147079, 29.8587554}, {106.4153027, 29.8580178}},
	{{106.4158024, 29.8573516}, {106.4162307, 29.8568043}},
	{{106.4166590, 29.8561380}, {106.4172777, 29.8554004}},
	{{106.4179677, 29.8542821}, {106.4185626, 29.8536634}},
	{{106.4194906, 29.8526505}, {106.4205185, 29.8516226}},
	{{106.4221175, 29.8496810}, {106.4231454, 29.8487102}},
	{{106.4252584, 29.8457406}, {106.4262863, 29.8446556}},
	{{106.4288562, 29.8420857}, {106.4299412, 29.8410578}},
	{{106.4329108, 29.8393446}, {106.4345669, 29.8386022}},
	{{106.4375364, 29.8375172}, {106.4391354, 29.8364892}},
	{{106.4412484, 29.8355184}, {106.4429616, 29.8344905}},
	{{106.4445606, 29.8331770}, {106.4458170, 29.8318065}},
	{{106.4471304, 29.8301504}, {106.4480441, 29.8288369}},
	{{106.4478157, 29.8256960}, {106.4470162, 29.8240970}},
	{{106.4469020, 29.8205564}, {106.4479299, 29.8193571}},
	{{106.4493576, 29.8175868}, {106.4503992, 29.8168673}},
	{{106.4527749, 29.8157708}, {106.4540541, 29.8152683}},
	{{106.4563384, 29.8138977}, {106.4575719, 29.8133495}},
	{{106.4598105, 29.8119789}, {106.4608156, 29.8112479}},
	{{106.4619577, 29.8097403}, {106.4629628, 29.8086895}},
	{{106.4644247, 29.8069078}, {106.4652927, 29.8060398}},
	{{106.4673943, 29.8036641}, {106.4685364, 29.8024306}},
	{{106.4699527, 29.8006945}, {106.4710491, 29.7997351}},
	{{106.4729679, 29.7981818}, {106.4737903, 29.7974509}},
	{{106.4755720, 29.7959432}, {106.4765771, 29.7949838}},
	{{106.4776736, 29.7932478}, {106.4785873, 29.7921513}},
	{{106.4802320, 29.7900498}, {106.4812827, 29.7891361}},
	{{106.4832929, 29.7873086}, {106.4842980, 29.7863492}},
	{{106.4860797, 29.7845218}, {106.4870391, 29.7836081}},
	{{106.4890950, 29.7817807}, {106.4898716, 29.7810497}},
	{{106.4920189, 29.7793593}, {106.4930696, 29.7786284}},
	{{106.4950341, 29.7772121}, {106.4963590, 29.7766639}},
	{{106.4983692, 29.7767553}, {106.4995570, 29.7777603}},
	{{106.5013730, 29.7777004}, {106.5023010, 29.7772721}},
	{{106.5054419, 29.7759634}, {106.5065126, 29.7754875}},
	{{106.5084162, 29.7742977}, {106.5089873, 29.7738456}},
	{{106.5104387, 29.7727987}, {106.5113429, 29.7722752}},
	{{106.5126754, 29.7712520}, {106.5133893, 29.7707286}},
	{{106.5145076, 29.7693009}, {106.5150549, 29.7685394}},
	{{106.5160305, 29.7663503}, {106.5162446, 29.7655413}},
	{{106.5166491, 29.7638043}, {106.5167681, 29.7626384}},
	{{106.5169823, 29.7606872}, {106.5170060, 29.7599020}},
	{{106.5170774, 29.7582364}, {106.5171012, 29.7574750}},
	{{106.5172202, 29.7554524}, {106.5173392, 29.7545244}},
	{{106.5180530, 29.7515025}, {106.5184099, 29.7505031}},
	{{106.5195045, 29.7482189}, {106.5199804, 29.7473622}},
	{{106.5213367, 29.7453873}, {106.5220743, 29.7446259}},
	{{106.5237161, 29.7434599}, {106.5245251, 29.7430316}},
	{{106.5260718, 29.7421512}, {106.5267380, 29.7415088}},
	{{106.5279992, 29.7402715}, {106.5284037, 29.7397480}},
	{{106.5294030, 29.7386058}, {106.5298075, 29.7381061}},
	{{106.5308069, 29.7367023}, {106.5312590, 29.7359408}},
	{{106.5319015, 29.7348939}, {106.5322584, 29.7342514}},
	{{106.5330436, 29.7329427}, {106.5332340, 29.7322527}},
	{{106.5334957, 29.7308964}, {106.5335433, 29.7299684}},
	{{106.5336623, 29.7283979}, {106.5334957, 29.7275175}},
	{{106.5333054, 29.7262326}, {106.5328057, 29.7253522}},
	{{106.5320918, 29.7242577}, {106.5314732, 29.7236628}},
	{{106.5303786, 29.7227348}, {106.5296886, 29.7222351}},
	{{106.5286654, 29.7215927}, {106.5278802, 29.7210454}},
	{{106.5268570, 29.7201650}, {106.5260242, 29.7196415}},
	{{106.5242634, 29.7188801}, {106.5233592, 29.7184280}},
	{{106.5217888, 29.7177142}, {106.5209322, 29.7173097}},
	{{106.5191258, 29.7163853}, {106.5176505, 29.7157191}},
	{{106.5159849, 29.7150528}, {106.5148903, 29.7144342}},
	{{106.5131771, 29.7134824}, {106.5120826, 29.7125306}},
	{{106.5105597, 29.7115788}, {106.5097031, 29.7104367}},
	{{106.5091796, 29.7089614}, {106.5092272, 29.7075813}},
	{{106.5094652, 29.7057254}, {106.5097507, 29.7043453}},
	{{106.5103694, 29.7025845}, {106.5107977, 29.7013471}},
	{{106.5113212, 29.6993960}, {106.5115591, 29.6980159}},
	{{106.5112736, 29.6961599}, {106.5107501, 29.6945419}},
	{{106.5096079, 29.6921148}, {106.5087037, 29.6907348}},
	{{106.5071809, 29.6889264}, {106.5056104, 29.6884505}},
	{{106.5024220, 29.6887360}, {106.5008991, 29.6894498}},
	{{106.4969016, 29.6908299}, {106.4951884, 29.6913058}},
	{{106.4920951, 29.6924004}, {106.4906674, 29.6940184}},
	{{106.4886211, 29.6973497}, {106.4876693, 29.6986822}},
	{{106.4847188, 29.7018230}, {106.4834339, 29.7023465}},
	{{106.4807213, 29.7029176}, {106.4785798, 29.7029652}},
	{{106.4761527, 29.7029652}, {106.4741540, 29.7028224}},
	{{106.4719173, 29.7025369}, {106.4700137, 29.7022038}},
	{{106.4663969, 29.7015851}, {106.4646361, 29.7010616}},
	{{106.4624946, 29.6999195}, {106.4614952, 29.6991105}},
	{{106.4594013, 29.6974448}, {106.4583068, 29.6966358}},
	{{106.4563556, 29.6948274}, {106.4552135, 29.6938281}},
	{{106.4537382, 29.6916865}, {106.4530719, 29.6906872}},
	{{106.4525009, 29.6888788}, {106.4521202, 29.6874987}},
	{{106.4518346, 29.6853572}, {106.4515967, 29.6841674}},
	{{106.4509780, 29.6820735}, {106.4507877, 29.6806934}},
	{{106.4504165, 29.6780316}, {106.4511018, 29.6755950}},
	{{106.3728584, 29.7434583}, {106.3725538, 29.7439722}},
	{{106.3717543, 29.7454570}, {106.3715068, 29.7460661}},
	{{106.3708977, 29.7474557}, {106.3708596, 29.7481220}},
	{{106.3711946, 29.7485522}, {106.3714002, 29.7486360}},
	{{106.3719485, 29.7486893}, {106.3721845, 29.7486816}},
	{{106.3728089, 29.7486055}, {106.3730830, 29.7485522}},
	{{106.3735931, 29.7483771}, {106.3738216, 29.7482629}},
	{{106.3741718, 29.7480268}, {106.3743622, 29.7478974}},
	{{106.3746896, 29.7475852}, {106.3748952, 29.7474405}},
	{{106.3753368, 29.7470598}, {106.3755043, 29.7469380}},
	{{106.3759155, 29.7466943}, {106.3761135, 29.7466486}},
	{{106.3765779, 29.7465649}, {106.3768292, 29.7466258}},
	{{106.3771338, 29.7467933}, {106.3772556, 29.7469684}},
	{{106.3775830, 29.7475243}, {106.3776896, 29.7478060}},
	{{106.3777886, 29.7482857}, {106.3778724, 29.7485294}},
	{{106.3779942, 29.7489481}, {106.3780627, 29.7492527}},
	{{106.3781541, 29.7500674}, {106.3780323, 29.7503187}},
	{{106.3772099, 29.7506081}, {106.3769510, 29.7506537}},
	{{106.3763800, 29.7509355}, {106.3762581, 29.7511258}},
	{{106.3758698, 29.7519177}, {106.3757708, 29.7521690}},
	{{106.3754891, 29.7528923}, {106.3754053, 29.7531817}},
	{{106.3753292, 29.7538594}, {106.3753597, 29.7541182}},
	{{106.3757099, 29.7546665}, {106.3759231, 29.7547807}},
	{{106.3769587, 29.7548568}, {106.3772632, 29.7549177}},
	{{106.3780780, 29.7550776}, {106.3783216, 29.7551614}},
	{{106.3791820, 29.7555269}, {106.3794105, 29.7556639}},
	{{106.3802404, 29.7561513}, {106.3805221, 29.7562655}},
	{{106.3814587, 29.7564787}, {106.3817633, 29.7565243}},
	{{106.3824638, 29.7565624}, {106.3827683, 29.7565929}},
	{{106.3835983, 29.7566005}, {106.3838724, 29.7566157}},
	{{106.3847861, 29.7567375}, {106.3850679, 29.7568137}},
	{{106.3857608, 29.7571411}, {106.3860273, 29.7573771}},
	{{106.3865983, 29.7579482}, {106.3867811, 29.7581614}},
	{{106.3872303, 29.7586944}, {106.3873826, 29.7589152}},
	{{106.3875730, 29.7594711}, {106.3876415, 29.7597604}},
	{{106.3878547, 29.7605142}, {106.3880450, 29.7607274}},
	{{106.3886846, 29.7609559}, {106.3889816, 29.7608721}},
	{{106.3897126, 29.7605523}, {106.3899562, 29.7604305}},
	{{106.3905121, 29.7601411}, {106.3907481, 29.7600574}},
	{{106.3913040, 29.7598061}, {106.3915400, 29.7598366}},
	{{106.3921720, 29.7602325}, {106.3924004, 29.7603619}},
	{{106.3932760, 29.7605980}, {106.3936187, 29.7606970}},
	{{106.3945172, 29.7608721}, {106.3947913, 29.7609939}},
	{{106.3955070, 29.7612376}, {106.3958116, 29.7614965}},
	{{106.3961238, 29.7621665}, {106.3962532, 29.7624939}},
	{{106.3971974, 29.7626310}, {106.3975933, 29.7626691}},
	{{106.3986670, 29.7628975}, {106.3988878, 29.7629508}},
	{{106.3997710, 29.7629280}, {106.3999690, 29.7627985}},
	{{106.4006238, 29.7621589}, {106.4008066, 29.7619609}},
	{{106.4016746, 29.7613442}, {106.4019259, 29.7611919}},
	{{106.4023979, 29.7610548}, {106.4026035, 29.7611005}},
	{{106.4029233, 29.7612680}, {106.4030756, 29.7615726}},
	{{106.4033878, 29.7624711}, {106.4034563, 29.7628137}},
	{{106.4037000, 29.7638874}, {106.4037761, 29.7641539}},
	{{106.4038827, 29.7648239}, {106.4039256, 29.7649791}},
	{{106.4040084, 29.7655872}, {106.4039941, 29.7657157}},
	{{106.4038094, 29.7664584}, {106.4037523, 29.7666440}},
	{{106.4035049, 29.7675053}, {106.4034573, 29.7677100}},
	{{106.4031546, 29.7686218}, {106.4030937, 29.7689568}},
	{{106.4029719, 29.7696649}, {106.4029490, 29.7699619}},
	{{106.4031165, 29.7705025}, {106.4032231, 29.7707385}},
	{{106.4035458, 29.7711925}, {106.4036648, 29.7713305}},
	{{106.4041216, 29.7717636}, {106.4042692, 29.7719064}},
	{{106.4048022, 29.7721919}, {106.4049830, 29.7722728}},
	{{106.4055160, 29.7724346}, {106.4057064, 29.7724774}},
	{{106.4061299, 29.7725821}, {106.4062679, 29.7726868}},
	{{106.4066201, 29.7731580}, {106.4067010, 29.7733341}},
	{{106.4068580, 29.7738433}, {106.4068913, 29.7740907}},
	{{106.4068628, 29.7748331}, {106.4068342, 29.7750282}},
	{{106.4067010, 29.7758420}, {106.4066867, 29.7760419}},
	{{106.4067771, 29.7765320}, {106.4069151, 29.7766272}},
	{{106.4074624, 29.7765844}, {106.4076480, 29.7765511}},
	{{106.4082143, 29.7763940}, {106.4084189, 29.7763702}},
	{{106.4088139, 29.7763227}, {106.4089377, 29.7764654}},
	{{106.4090709, 29.7768699}, {106.4090947, 29.7770841}},
	{{106.4090343, 29.7777602}, {106.4090248, 29.7780267}},
	{{106.4088249, 29.7785597}, {106.4086726, 29.7788071}},
	{{106.4083585, 29.7792545}, {106.4081206, 29.7795019}},
	{{106.4073591, 29.7803966}, {106.4073020, 29.7806536}},
	{{106.4076828, 29.7811580}, {106.4080349, 29.7813389}},
	{{106.4087488, 29.7816244}, {106.4090343, 29.7817481}},
	{{106.4096910, 29.7820146}, {106.4099670, 29.7822050}},
	{{106.4105381, 29.7825286}, {106.4107570, 29.7826904}},
	{{106.4112995, 29.7831758}, {106.4114899, 29.7834328}},
	{{106.4118992, 29.7841561}, {106.4118896, 29.7844893}},
	{{106.4115946, 29.7851365}, {106.4113376, 29.7853744}},
	{{106.4104905, 29.7859455}, {106.4101669, 29.7861168}},
	{{106.4092913, 29.7867736}, {106.4091199, 29.7870781}},
	{{106.4087678, 29.7882583}, {106.4087678, 29.7885534}},
	{{106.4088344, 29.7894957}, {106.4088534, 29.7897907}},
	{{106.4084347, 29.7906378}, {106.4081872, 29.7909138}},
	{{106.4077589, 29.7915801}, {106.4076447, 29.7918942}},
	{{106.4076637, 29.7928079}, {106.4077970, 29.7930553}},
	{{106.4084156, 29.7936645}, {106.4086536, 29.7938739}},
	{{106.4093674, 29.7943402}, {106.4096054, 29.7945306}},
	{{106.4105667, 29.7951207}, {106.4108332, 29.7953301}},
	{{106.4112044, 29.7956823}, {106.4113091, 29.7959678}},
	{{106.4113852, 29.7969005}, {106.4112234, 29.7971861}},
	{{106.4104049, 29.7977191}, {106.4101193, 29.7978714}},
	{{106.4097767, 29.7983473}, {106.4098814, 29.7987470}},
	{{106.4104715, 29.7992229}, {106.4107570, 29.7994418}},
	{{106.4115755, 29.7998796}, {106.4119943, 29.8000700}},
	{{106.4130565, 29.8002394}, {106.4133306, 29.8001994}},
	{{106.4142672, 29.7999881}, {106.4144671, 29.7999596}},
	{{106.4151695, 29.7997940}, {106.4153808, 29.7997711}},
	{{106.4159861, 29.7998853}, {106.4160889, 29.8000053}},
	{{106.4160946, 29.8007762}, {106.4159576, 29.8010103}},
	{{106.4153694, 29.8017984}, {106.4152152, 29.8019640}},
	{{106.4145870, 29.8025294}, {106.4144157, 29.8026436}},
	{{106.4137761, 29.8029920}, {106.4135362, 29.8030434}},
	{{106.4129423, 29.8031861}, {106.4127596, 29.8032318}},
	{{106.4122799, 29.8034088}, {106.4121828, 29.8035630}},
	{{106.4121428, 29.8039856}, {106.4122570, 29.8041912}},
	{{106.4128452, 29.8044425}, {106.4131022, 29.8044882}},
	{{106.4139360, 29.8045224}, {106.4142329, 29.8044939}},
	{{106.4150439, 29.8042826}, {106.4152323, 29.8042312}},
	{{106.4158605, 29.8039285}, {106.4160546, 29.8038371}},
	{{106.4166314, 29.8034203}, {106.4168027, 29.8033003}},
	{{106.4172596, 29.8028606}, {106.4173852, 29.8027521}},
	{{106.4178592, 29.8024209}, {106.4180648, 29.8023124}},
	{{106.4185616, 29.8020668}, {106.4187958, 29.8019755}},
	{{106.4193097, 29.8018441}, {106.4194639, 29.8019355}},
	{{106.4196524, 29.8023752}, {106.4196924, 29.8026093}},
	{{106.4196067, 29.8031804}, {106.4195096, 29.8034260}},
	{{106.4190528, 29.8039456}, {106.4188929, 29.8040770}},
	{{106.4184246, 29.8043739}, {106.4181962, 29.8045053}},
	{{106.4176651, 29.8048137}, {106.4174538, 29.8049279}},
	{{106.4169341, 29.8052420}, {106.4167628, 29.8053619}},
	{{106.4162145, 29.8056703}, {106.4160261, 29.8058245}},
	{{106.4154836, 29.8062585}, {106.4153180, 29.8064127}},
	{{106.4149468, 29.8068924}, {106.4148668, 29.8071094}},
	{{106.4151409, 29.8077890}, {106.4152551, 29.8079774}},
	{{106.4157063, 29.8083372}, {106.4159119, 29.8084286}},
	{{106.4168085, 29.8084571}, {106.4171111, 29.8084457}},
	{{106.4179392, 29.8082401}, {106.4181905, 29.8081944}},
	{{106.4189443, 29.8079945}, {106.4191441, 29.8079489}},
	{{106.4196695, 29.8078118}, {106.4198466, 29.8077547}},
	{{106.4203320, 29.8075662}, {106.4204919, 29.8075034}},
	{{106.4210058, 29.8072921}, {106.4211657, 29.8072464}},
	{{106.4216968, 29.8070808}, {106.4218853, 29.8070351}},
	{{106.4224278, 29.8070523}, {106.4225991, 29.8071208}},
	{{106.4229189, 29.8073378}, {106.4230388, 29.8074863}},
	{{106.4233815, 29.8080459}, {106.4234557, 29.8082058}},
	{{106.4236156, 29.8087198}, {106.4236042, 29.8088683}},
	{{106.4233472, 29.8092737}, {106.4231702, 29.8094279}},
	{{106.4223307, 29.8099590}, {106.4221423, 29.8100789}},
	{{106.4214627, 29.8105015}, {106.4213142, 29.8106272}},
	{{106.4210972, 29.8111982}, {106.4211200, 29.8114552}},
	{{106.4213256, 29.8119635}, {106.4213999, 29.8121691}},
	{{106.4216055, 29.8126716}, {106.4216226, 29.8129115}},
	{{106.4215940, 29.8135225}, {106.4215712, 29.8137338}},
	{{106.4214341, 29.8143962}, {106.4213999, 29.8145733}},
	{{106.4212914, 29.8150644}, {106.4212971, 29.8152643}},
	{{106.4213371, 29.8157268}, {106.4213999, 29.8158867}},
	{{106.4216397, 29.8161437}, {106.4217539, 29.8163265}},
	{{106.4220794, 29.8167091}, {106.4221651, 29.8168461}},
	{{106.4224335, 29.8172516}, {106.4225477, 29.8174115}},
	{{106.4227761, 29.8178341}, {106.4228618, 29.8180168}},
	{{106.4230503, 29.8183081}, {106.4231816, 29.8185079}},
	{{106.4235128, 29.8188563}, {106.4236556, 29.8189876}},
	{{106.4241467, 29.8193646}, {106.4243180, 29.8194731}},
	{{106.4251518, 29.8196044}, {106.4253574, 29.8196272}},
	{{106.4261797, 29.8196387}, {106.4264139, 29.8196615}},
	{{106.4271734, 29.8197529}, {106.4273904, 29.8198157}},
	{{106.4278301, 29.8198728}, {106.4280757, 29.8198842}},
	{{106.4286982, 29.8198956}, {106.4288752, 29.8198328}},
	{{106.4293891, 29.8195359}, {106.4295319, 29.8194673}},
	{{106.4301715, 29.8190448}, {106.4304228, 29.8190105}},
	{{106.4310110, 29.8189477}, {106.4312623, 29.8189534}},
	{{106.4318847, 29.8189420}, {106.4321189, 29.8189591}},
	{{106.4326956, 29.8189077}, {106.4328555, 29.8188791}},
	{{106.4331468, 29.8187021}, {106.4332439, 29.8185651}},
	{{106.4332953, 29.8182395}, {106.4332553, 29.8180397}},
	{{106.4330726, 29.8176970}, {106.4329640, 29.8175485}},
	{{106.4327585, 29.8172516}, {106.4326500, 29.8171088}},
	{{106.4324958, 29.8168176}, {106.4324272, 29.8166634}},
	{{106.4322788, 29.8162751}, {106.4323473, 29.8161095}},
	{{106.4325871, 29.8158410}, {106.4327699, 29.8158239}},
	{{106.4331525, 29.8160009}, {106.4333238, 29.8161494}},
	{{106.4339234, 29.8165892}, {106.4341005, 29.8167148}},
	{{106.4346773, 29.8170003}, {106.4349057, 29.8170917}},
	{{106.4354710, 29.8171774}, {106.4357395, 29.8172059}},
	{{106.4364304, 29.8171317}, {106.4366817, 29.8170803}},
	{{106.4372185, 29.8170346}, {106.4374298, 29.8170232}},
	{{106.4379495, 29.8169718}, {106.4381665, 29.8169661}},
	{{106.4386633, 29.8170117}, {106.4388404, 29.8170860}},
	{{106.4391545, 29.8173544}, {106.4392515, 29.8175543}},
	{{106.4393315, 29.8180397}, {106.4393258, 29.8182738}},
	{{106.4393372, 29.8188220}, {106.4393829, 29.8190276}},
	{{106.4396342, 29.8194788}, {106.4397427, 29.8196330}},
	{{106.4399882, 29.8199014}, {106.4401081, 29.8200213}},
	{{106.4403480, 29.8202212}, {106.4404965, 29.8203068}},
	{{106.4409076, 29.8204496}, {106.4410790, 29.8204896}},
	{{106.4415530, 29.8205124}, {106.4417814, 29.8205124}},
	{{106.4423639, 29.8204382}, {106.4426380, 29.8204096}},
	{{106.4431577, 29.8203011}, {106.4433575, 29.8202840}},
	{{106.4437744, 29.8202611}, {106.4439914, 29.8202497}},
}

func TestName(t *testing.T) {
	for _, v := range data {
		p, angle := GetCenterAndAngle(v)
		fmt.Printf("[%.7f,%.7f] %.3f\n", p[0], p[1], angle)
	}
}