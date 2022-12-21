package speedtestnet

func GetServerByID(id string) server {
	servers := getServerList()

	for _, nextServer := range servers {
		if nextServer.ID == id {
			return nextServer
		}
	}

	return server{}
}

func getServerList() []server {

	return []server{
		{
			URL:       "http://speedtest.aptalaska.net/speedtest/upload.php",
			Latitude:  58.3833,
			Longitude: -134.1833,
			Name:      "Juneau, AK",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "Alaska Power and Telephone",
			ID:        "10367",
			URL2:      "http://speedtest2.aptalaska.net/speedtest/upload.php",
			Host:      "speedtest.aptalaska.net:8080",
		},
		{
			URL:       "http://caro.OST.myvzw.com/speedtest/upload.php",
			Latitude:  38.7907,
			Longitude: -121.2358,
			Name:      "Rocklin, CA",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "Verizon",
			ID:        "14810",
			Host:      "caro.OST.myvzw.com:8080",
		},
		{
			URL:       "http://speedtest01.serverpronto.com/upload.php",
			Latitude:  25.7878,
			Longitude: -80.2242,
			Name:      "Miami, FL",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "ServerPronto",
			ID:        "12215",
			Host:      "Serverpronto.com:8080",
		},
		{
			URL:       "http://orl.speedtest.t-mobile.com/speedtest/upload.jsp",
			Latitude:  28.4158,
			Longitude: -81.2989,
			Name:      "Orlando, FL",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "T-Mobile",
			ID:        "1762",
			Host:      "orl.speedtest.t-mobile.com:8080",
		},
		{
			URL:       "http://speedtest.cityofblakely.net/speedtest/upload.php",
			Latitude:  31.3777,
			Longitude: -84.9341,
			Name:      "Blakely, GA",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "City of Blakely",
			ID:        "12395",
			Host:      "speedtest.cityofblakely.net:8080",
		},
		{
			URL:       "http://pqi.pwless.net/speedtest/upload.php",
			Latitude:  46.6793,
			Longitude: -68.0022,
			Name:      "Presque Isle, ME",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "Pioneer Wireless",
			ID:        "534",
			URL2:      "http://pqi2.pwless.net/speedtest/upload.php",
			Host:      "pqi.pwless.net:8080",
		},
		{
			URL:       "http://jcy1sp1.qtsdatacenters.com/speedtest/upload.php",
			Latitude:  40.7282,
			Longitude: -74.0776,
			Name:      "Jersey City, NJ",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "QTS TaskData Centers",
			ID:        "14854",
			Host:      "jcy1sp1.qtsdatacenters.com:8080",
		},
		{
			URL:       "http://speedtest.nyc.rr.com/speedtest/upload.php",
			Latitude:  40.7127,
			Longitude: -74.0059,
			Name:      "New York City, NY",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "Spectrum",
			ID:        "16976",
			Host:      "speedtest.nyc.rr.com:8080",
		},
		{
			URL:       "http://speedtest.gethotwired.com/speedtest/upload.php",
			Latitude:  39.9500,
			Longitude: -75.1667,
			Name:      "Philadelphia, PA",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "Hotwire Fision",
			ID:        "4847",
			URL2:      "http://speedtestpa.gethotwired.com/speedtest/upload.php",
			Host:      "speedtest.gethotwired.com:8080",
		},
		{
			URL:       "http://hou.speedtest.t-mobile.com/speedtest/upload.jsp",
			Latitude:  29.7631,
			Longitude: -95.3631,
			Name:      "Houston, TX",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "T-Mobile",
			ID:        "1816",
			Host:      "hou.speedtest.t-mobile.com:8080",
		},
		{
			URL:       "http://stosat-rstn-01.sys.comcast.net/speedtest/upload.php",
			Latitude:  38.7514,
			Longitude: -77.4764,
			Name:      "Manassas, VA",
			Country:   "United States",
			CC:        "US",
			Sponsor:   "Comcast",
			ID:        "9839",
			URL2:      "http://a-stosat-rstn-01.sys.comcast.net/speedtest/upload.php",
			Host:      "stosat-rstn-01.sys.comcast.net:8080",
		},
		{
			URL:       "http://murmansk.speedtest.rt.ru/speedtest/upload.php",
			Latitude:  68.9543,
			Longitude: 33.1076,
			Name:      "Murmansk",
			Country:   "Russian Federation",
			CC:        "RU",
			Sponsor:   "Rostelecom",
			ID:        "3522",
			URL2:      "http://murmansk.speedtest2.rt.ru/speedtest/upload.php",
			Host:      "murmansk.speedtest.rt.ru:8080",
		},
		{
			URL:       "http://speedtest.netplaza.fi/speedtest/upload.php",
			Latitude:  65.0167,
			Longitude: 25.4667,
			Name:      "Oulu",
			Country:   "Finland",
			CC:        "FI",
			Sponsor:   "Netplaza Oy",
			ID:        "2138",
			URL2:      "http://speedtest2.netplaza.fi/speedtest/upload.php",
			Host:      "speedtest.netplaza.fi:8080",
		},
		{
			URL:       "http://speedtest.c.is/speedtest/upload.php",
			Latitude:  64.1333,
			Longitude: -21.9333,
			Name:      "Reykjavík",
			Country:   "Iceland",
			CC:        "IS",
			Sponsor:   "Vodafone",
			ID:        "4141",
			URL2:      "http://speedtest2.c.is/speedtest/upload.php",
			Host:      "speedtest.c.is:8080",
		},
		{
			URL:       "http://ookla-umea.alltele.net/upload.php",
			Latitude:  63.8258,
			Longitude: 20.2630,
			Name:      "Umea",
			Country:   "Sweden",
			CC:        "SE",
			Sponsor:   "AllTele AB",
			ID:        "10928",
			URL2:      "http://ookla-umea.t3.se/upload.php",
			Host:      "ookla-umea.alltele.net:8080",
		},
		{
			URL:       "http://speedtest-sg.karelia.pro/speedtest/upload.php",
			Latitude:  63.7388,
			Longitude: 34.3098,
			Name:      "Segezha",
			Country:   "Russian Federation",
			CC:        "RU",
			Sponsor:   "CityLink Ltd",
			ID:        "8657",
			URL2:      "http://speedtest-sg2.karelia.pro/speedtest/upload.php",
			Host:      "speedtest-sg.karelia.pro:8080",
		},
		{
			URL:       "http://yk-netspeed.nwtel.ca/speedtest/upload.php",
			Latitude:  62.4422,
			Longitude: -114.3975,
			Name:      "Yellowknife, NT",
			Country:   "Canada",
			CC:        "CA",
			Sponsor:   "Northwestel Inc.",
			ID:        "13139",
			Host:      "yk-netspeed.nwtel.ca:8080",
		},
		{
			URL:       "http://speedtest1.internet.fo/speedtest/upload.php",
			Latitude:  62.0000,
			Longitude: -6.7833,
			Name:      "Torshavn",
			Country:   "Faroe Islands",
			CC:        "FO",
			Sponsor:   "Faroese Telecom",
			ID:        "462",
			URL2:      "http://speedtest2.internet.fo/speedtest/upload.php",
			Host:      "speedtest1.internet.fo:8080",
		},
		{
			URL:       "http://speedtest1.saunalahti.fi/speedtest/upload.php",
			Latitude:  60.1708,
			Longitude: 24.9375,
			Name:      "Helsinki",
			Country:   "Finland",
			CC:        "FI",
			Sponsor:   "Elisa Oyj",
			ID:        "4549",
			URL2:      "http://speedtest1-1.saunalahti.fi/speedtest/upload.php",
			Host:      "speedtest1.saunalahti.fi:8080",
		},
		{
			URL:       "http://speedtest1.telenor.net/mini/speedtest/upload.php",
			Latitude:  59.9494,
			Longitude: 10.7564,
			Name:      "Oslo",
			Country:   "Norway",
			CC:        "NO",
			Sponsor:   "Telenor Norge AS",
			ID:        "3672",
			Host:      "speedtest1.telenor.net:8080",
		},
		{
			URL:       "http://sp.runnet.ru/speedtest/upload.php",
			Latitude:  59.9343,
			Longitude: 30.3351,
			Name:      "Saint Petersburg",
			Country:   "Russian Federation",
			CC:        "RU",
			Sponsor:   "RUNNet",
			ID:        "14633",
			Host:      "sp.runnet.ru:8080",
		},
		{
			URL:       "http://s1.sto.speedtest.interoute.net/speedtest/upload.php",
			Latitude:  59.3294,
			Longitude: 18.0686,
			Name:      "Stockholm",
			Country:   "Sweden",
			CC:        "SE",
			Sponsor:   "Interoute VDC",
			ID:        "10256",
			URL2:      "http://s2.sto.speedtest.interoute.net/speedtest/upload.php",
			Host:      "sto.speedtest.interoute.net:8080",
		},
		{
			URL:       "http://speedtest.nygreenit.net/speedtest/upload.php",
			Latitude:  57.5833,
			Longitude: 9.9500,
			Name:      "Hirtshals",
			Country:   "Denmark",
			CC:        "DK",
			Sponsor:   "Nygreen IT ApS",
			ID:        "5294",
			URL2:      "http://speedtest.nygreenit.com/speedtest/upload.php",
			Host:      "speedtest.nygreenit.net:8080",
		},
		{
			URL:       "http://speedtest1.converged.co.uk/upload.php",
			Latitude:  57.1526,
			Longitude: -2.1100,
			Name:      "Aberdeen",
			Country:   "Great Britain",
			CC:        "GB",
			Sponsor:   "Converged Communication Solutions Limited",
			ID:        "9171",
			URL2:      "http://speedtest2.converged.co.uk/upload.php",
			Host:      "speedtest1.converged.co.uk:8080",
		},
		{
			URL:       "http://ekt1.companion.tele2.ru/speedtest/upload.php",
			Latitude:  56.8333,
			Longitude: 60.5833,
			Name:      "Ekaterinburg",
			Country:   "Russian Federation",
			CC:        "RU",
			Sponsor:   "Tele2 Russia",
			ID:        "6616",
			URL2:      "http://176.59.223.150/speedtest/upload.php",
			Host:      "ekt1.companion.tele2.ru:8080",
		},
		{
			URL:       "http://moscow.speedtest.rt.ru/speedtest/upload.php",
			Latitude:  55.7517,
			Longitude: 37.6178,
			Name:      "Moscow",
			Country:   "Russian Federation",
			CC:        "RU",
			Sponsor:   "Rostelecom",
			ID:        "3682",
			URL2:      "http://moscow.speedtest2.rt.ru/speedtest/upload.php",
			Host:      "moscow.speedtest.rt.ru:8080",
		},
		{
			URL:       "http://ptr-st1.2day.kz/speedtest/upload.php",
			Latitude:  54.8833,
			Longitude: 69.1667,
			Name:      "Petropavlovsk",
			Country:   "Kazakhstan",
			CC:        "KZ",
			Sponsor:   "Beeline KZ",
			ID:        "6971",
			URL2:      "http://ptr-st2.2day.kz/speedtest/upload.php",
			Host:      "ptr-st1.2day.kz:8080",
		},
		{
			URL:       "http://speed3.plastcom.pl/speedtest/upload.php",
			Latitude:  54.6167,
			Longitude: 18.3500,
			Name:      "Reda",
			Country:   "Poland",
			CC:        "PL",
			Sponsor:   "Plast-Com",
			ID:        "4216",
			URL2:      "http://speed4.plastcom.pl/speedtest/upload.php",
			Host:      "speed3.plastcom.pl:8080",
		},
		{
			URL:       "http://4gnanjing1.speedtest.jsinfo.net/speedtest/upload.jsp",
			Latitude:  32.0500,
			Longitude: 118.7667,
			Name:      "Nanjing",
			Country:   "China",
			CC:        "CN",
			Sponsor:   "China Telecom JiangSu Branch",
			ID:        "5316",
			Host:      "4gnanjing1.speedtest.jsinfo.net:8080",
		},
		{
			URL:       "http://ouarglaspeed.djaweb.dz/speedtest/upload.php",
			Latitude:  31.9500,
			Longitude: 5.3167,
			Name:      "Ouargla",
			Country:   "Algeria",
			CC:        "DZ",
			Sponsor:   "Algeria Telecom",
			ID:        "5848",
			Host:      "ouarglaspeed.djaweb.dz:8080",
		},
		{
			URL:       "http://speed1.jcs.jo/speedtest/upload.php",
			Latitude:  31.9500,
			Longitude: 35.9333,
			Name:      "Amman",
			Country:   "Jordan",
			CC:        "JO",
			Sponsor:   "Jordanian Cable Services",
			ID:        "7542",
			URL2:      "http://speed2.jcs.jo/speedtest/upload.php",
			Host:      "speed1.jcs.jo:8080",
		},
		{
			URL:       "http://speedtest-rm.paltel.ps/speedtest/upload.php",
			Latitude:  31.9000,
			Longitude: 35.2000,
			Name:      "Ramallah",
			Country:   "Palestine",
			CC:        "PS",
			Sponsor:   "PALTEL Group",
			ID:        "1294",
			URL2:      "http://82.213.28.12/speedtest/upload.php",
			Host:      "82.213.28.12:8080",
		},
		{
			URL:       "http://speedtest.coolnet.ps/speedtest/upload.aspx",
			Latitude:  31.7833,
			Longitude: 35.2167,
			Name:      "Jerusalem",
			Country:   "Israel",
			CC:        "IL",
			Sponsor:   "Coolnet",
			ID:        "1075",
			URL2:      "http://speedtest1.coolnet.ps/speedtest/upload.aspx",
			Host:      "speedtest.coolnet.ps:8080",
		},
		{
			URL:       "http://speedtest-marrakech1.hostoweb.com/speedtest/upload.php",
			Latitude:  31.6300,
			Longitude: -8.0089,
			Name:      "Marrakesh",
			Country:   "Morocco",
			CC:        "MA",
			Sponsor:   "HOSTOWEB.COM",
			ID:        "7574",
			URL2:      "http://speedtest-marrakech2.hostoweb.com/speedtest/upload.php",
			Host:      "speedtest-marrakech1.hostoweb.com:8080",
		},
		{
			URL:       "http://speedtest5.telenor.com.pk/speedtest/upload.php",
			Latitude:  31.5758,
			Longitude: 74.3269,
			Name:      "Lahore",
			Country:   "Pakistan",
			CC:        "PK",
			Sponsor:   "Telenor Pakistan",
			ID:        "9716",
			URL2:      "http://speedtest6.telenor.com.pk/speedtest/upload.php",
			Host:      "speedtest5.telenor.com.pk:8080",
		},
		{
			URL:       "http://201.174.27.254/speedtest/upload.php",
			Latitude:  31.3403,
			Longitude: -110.9336,
			Name:      "Nogales",
			Country:   "Mexico",
			CC:        "MX",
			Sponsor:   "TRANSTELCO",
			ID:        "3152",
			URL2:      "http://201.174.27.254/speedtest2/upload.php",
			Host:      "201.174.27.254:8080",
		},
		{
			URL:       "http://sp1.datawaveitsolutions.com/speedtest/upload.php",
			Latitude:  31.3260,
			Longitude: 75.5762,
			Name:      "Jalandhar",
			Country:   "India",
			CC:        "IN",
			Sponsor:   "DATAWAVE IT SOLUTIONS PVT LTD",
			ID:        "14842",
			Host:      "sp1.datawaveitsolutions.com:8080",
		},
		{
			URL:       "http://ahvaz1.irancell.ir/speedtest/upload.php",
			Latitude:  31.3203,
			Longitude: 48.6692,
			Name:      "Ahvaz",
			Country:   "Iran",
			CC:        "IR",
			Sponsor:   "MTNIrancell",
			ID:        "10272",
			URL2:      "http://ahvaz2.irancell.ir/speedtest/upload.php",
			Host:      "ahvaz1.irancell.ir:8080",
		},
		{
			URL:       "http://211.95.17.50/shunicom1/speedtest/upload.aspx",
			Latitude:  31.2000,
			Longitude: 121.5000,
			Name:      "Shanghai",
			Country:   "China",
			CC:        "CN",
			Sponsor:   "Shanghai Branch, China Unicom",
			ID:        "5083",
			URL2:      "http://211.95.17.50/shunicom2/speedtest/upload.php",
			Host:      "211.95.17.50:8080",
		},
		{
			URL:       "http://sp30.etisalatdata.net/speedtest/upload.php",
			Latitude:  31.1980,
			Longitude: 29.9192,
			Name:      "Alexandria",
			Country:   "Egypt",
			CC:        "EG",
			Sponsor:   "Etisalat Misr",
			ID:        "15495",
			Host:      "sp30.etisalatdata.net:8080",
		},
		{
			URL:       "http://speedshm1.jioconnect.com/speedtest/upload.php",
			Latitude:  31.1048,
			Longitude: 77.1734,
			Name:      "Shimla",
			Country:   "India",
			CC:        "IN",
			Sponsor:   "Jio",
			ID:        "10109",
			Host:      "speedshm1.jioconnect.com:8080",
		},
		{
			URL:       "http://sp1.uestc.edu.cn/upload.php",
			Latitude:  30.6586,
			Longitude: 104.0647,
			Name:      "Chengdu",
			Country:   "China",
			CC:        "CN",
			Sponsor:   "University of Electronic Science and Technology of China",
			ID:        "11444",
			Host:      "sp1.uestc.edu.cn:8080",
		},
		{
			URL:       "http://sp5.fastiraq.com/speedtest/upload.php",
			Latitude:  30.4944,
			Longitude: 47.8236,
			Name:      "Al Basrah",
			Country:   "Iraq",
			CC:        "IQ",
			Sponsor:   "Fastiraq LLC",
			ID:        "13380",
			Host:      "sp5.fastiraq.com:8080",
		},
		{
			URL:       "http://41.128.129.74/speedtest/upload.php",
			Latitude:  30.0566,
			Longitude: 31.2262,
			Name:      "Cairo",
			Country:   "Egypt",
			CC:        "EG",
			Sponsor:   "Orange DSL",
			ID:        "7054",
			Host:      "41.128.129.74:8080",
		},
		{
			URL:       "http://stc.qualitynet.net/upload.php",
			Latitude:  29.3544,
			Longitude: 48.0034,
			Name:      "Kuwait City",
			Country:   "Kuwait",
			CC:        "KW",
			Sponsor:   "Qualitynet",
			ID:        "14988",
			Host:      "stc.qualitynet.net:8080",
		},
		{
			URL:       "http://speedtest.gtel.in/upload.php",
			Latitude:  28.6100,
			Longitude: 77.2300,
			Name:      "New Delhi",
			Country:   "India",
			CC:        "IN",
			Sponsor:   "Gigantic Infotel Pvt Ltd",
			ID:        "8906",
			URL2:      "http://sp2.gtel.in/upload.php",
			Host:      "speedtest.gtel.in:8080",
		},
		{
			URL:       "http://st2.wificanarias.com/st/upload.php",
			Latitude:  28.3914,
			Longitude: -16.5239,
			Name:      "La Orotava",
			Country:   "Spain",
			CC:        "ES",
			Sponsor:   "WiFi Canarias",
			ID:        "11990",
			Host:      "st2.wificanarias.com:8080",
		},
		{
			URL:       "http://speedtest-xg.glbb.ne.jp/speedtest/upload.php",
			Latitude:  26.3167,
			Longitude: 127.7667,
			Name:      "Chatan",
			Country:   "Japan",
			CC:        "JP",
			Sponsor:   "GLBB Japan KK",
			ID:        "811",
			URL2:      "http://speedtest-xg.glbb.ne.jp/speedtest2/upload.php",
			Host:      "speedtest-xg.glbb.ne.jp:8080",
		},
		{
			URL:       "http://dam.myspeed.net.sa/speedtest/upload.php",
			Latitude:  26.2833,
			Longitude: 50.2000,
			Name:      "Dammam",
			Country:   "Saudi Arabia",
			CC:        "SA",
			Sponsor:   "Mobily",
			ID:        "1734",
			URL2:      "http://86.51.184.166/speedtest/upload.php",
			Host:      "dam.myspeed.net.sa:8080",
		},
		{
			URL:       "http://speedtest.viva.com.bh/speedtest/upload.php",
			Latitude:  26.2361,
			Longitude: 50.5831,
			Name:      "Seef",
			Country:   "Bahrain",
			CC:        "BH",
			Sponsor:   "Viva Bahrain",
			ID:        "1807",
			URL2:      "http://speedtest2.viva.com.bh/speedtest/upload.php",
			Host:      "speedtest.viva.com.bh:8080",
		},
		{
			URL:       "http://ookla.securehost.com/speedtest/upload.php",
			Latitude:  25.0600,
			Longitude: -77.3450,
			Name:      "Nassau",
			Country:   "Bahamas",
			CC:        "BS",
			Sponsor:   "Securehosting Ltd.",
			ID:        "6442",
			URL2:      "http://ookla1.securehost.com/speedtest/upload.php",
			Host:      "ookla.securehost.com:8080",
		},
		{
			URL:       "http://speedtest.taifo.com.tw/speedtest/upload.php",
			Latitude:  25.0374,
			Longitude: 121.5635,
			Name:      "Taipei",
			Country:   "Taiwan",
			CC:        "TW",
			Sponsor:   "TAIFO Taiwan",
			ID:        "13506",
			Host:      "speedtest.taifo.com.tw:8080",
		},
		{
			URL:       "http://ftp.qubee.com.bd/speedtest/upload.php",
			Latitude:  23.7000,
			Longitude: 90.3667,
			Name:      "Dhaka",
			Country:   "Bangladesh",
			CC:        "BD",
			Sponsor:   "Augere Wireless Broadband Bangladesh Ltd.",
			ID:        "4773",
			URL2:      "http://speedtest.qubee.com.bd/speedtest/upload.php",
			Host:      "ftp.qubee.com.bd:8080",
		},
		{
			URL:       "http://speedtest.awasr.com/speedtest/upload.php",
			Latitude:  23.6100,
			Longitude: 58.5400,
			Name:      "Muscat",
			Country:   "Oman",
			CC:        "OM",
			Sponsor:   "AWASR",
			ID:        "10477",
			URL2:      "http://speedtest1.awasr.com/speedtest/upload.php",
			Host:      "speedtest.awasr.com:8080",
		},
		{
			URL:       "http://ookla-speedtest.hgconair.hgc.com.hk/speedtest/upload.php",
			Latitude:  22.3771,
			Longitude: 114.1974,
			Name:      "Shatin",
			Country:   "Hong Kong",
			CC:        "HK",
			Sponsor:   "HGC Global Communications Limited",
			ID:        "16176",
			Host:      "ookla-speedtest.hgconair.hgc.com.hk:8080",
		},
		{
			URL:       "http://hnispeedtest.cmctelecom.vn/speedtest/upload.php",
			Latitude:  21.0285,
			Longitude: 105.8542,
			Name:      "Hanoi",
			Country:   "Vietnam",
			CC:        "VN",
			Sponsor:   "CMC Telecom",
			ID:        "6342",
			URL2:      "http://hnispeedtest.cmcti.vn/speedtest/upload.php",
			Host:      "hnispeedtest.cmctelecom.vn:8080",
		},
		{
			URL:       "http://speedtest05.spt.vn/upload.php",
			Latitude:  21.0285,
			Longitude: 105.8542,
			Name:      "Hanoi",
			Country:   "Vietnam",
			CC:        "VN",
			Sponsor:   "SAIGON POSTEL CORP.",
			ID:        "7215",
			URL2:      "http://speedtest06.spt.vn/upload.php",
			Host:      "speedtest05.spt.vn:8080",
		},
		{
			URL:       "http://cmi1.speedtest.trueinternet.co.th/speedtest/upload.php",
			Latitude:  18.7061,
			Longitude: 98.9817,
			Name:      "Chiang Mai",
			Country:   "Thailand",
			CC:        "TH",
			Sponsor:   "True Internet Corporation",
			ID:        "16396",
			Host:      "cmi1.speedtest.trueinternet.co.th:8080",
		},
		{
			URL:       "http://sbw1-test.nnpr.net/speedtest/upload.php",
			Latitude:  18.4517,
			Longitude: -66.0689,
			Name:      "San Juan",
			Country:   "Puerto Rico",
			CC:        "PR",
			Sponsor:   "Neptunomedia, Inc.",
			ID:        "2383",
			URL2:      "http://204.15.149.245/speedtest/upload.php",
			Host:      "sbw1-test.nnpr.net:8080",
		},
		{
			URL:       "http://speedtest.mmspeednet.com/speedtest/upload.php",
			Latitude:  16.8000,
			Longitude: 96.1500,
			Name:      "Yangon",
			Country:   "Myanmar",
			CC:        "MM",
			Sponsor:   "Myanmar Speed Net Co., Ltd",
			ID:        "12718",
			Host:      "speedtest.mmspeednet.com:8080",
		},
		{
			URL:       "http://www.tevisat.net/mini/speedtest/upload.aspx",
			Latitude:  15.7667,
			Longitude: -86.8333,
			Name:      "La Ceiba",
			Country:   "Honduras",
			CC:        "HN",
			Sponsor:   "Tevisat, S.A.",
			ID:        "4074",
			URL2:      "http://speedtest.tevisat.net/speedtest/upload.aspx",
			Host:      "www.tevisat.net:8080",
		},
		{
			URL:       "http://testerapido.cvmovel.cv/speedtest/upload.php",
			Latitude:  14.9150,
			Longitude: -23.5119,
			Name:      "Praia",
			Country:   "Cape Verde",
			CC:        "CV",
			Sponsor:   "CVMultimedia",
			ID:        "14005",
			Host:      "testerapido.cvmovel.cv:8080",
		},
		{
			URL:       "http://speedtest.tigo.sn/speedtest/upload.php",
			Latitude:  14.6962,
			Longitude: -17.4442,
			Name:      "Dakar",
			Country:   "Senegal",
			CC:        "SN",
			Sponsor:   "TIGO",
			ID:        "4331",
			URL2:      "http://speedtest2.tigo.sn/speedtest/upload.php",
			Host:      "speedtest.tigo.sn:8080",
		},
		{
			URL:       "http://greenhills.smart.com.ph/speedtest/upload.php",
			Latitude:  14.5800,
			Longitude: 121.0000,
			Name:      "Manila",
			Country:   "Philippines",
			CC:        "PH",
			Sponsor:   "Smart Communications Inc.",
			ID:        "7415",
			URL2:      "http://greenhills2.smart.com.ph/speedtest/upload.php",
			Host:      "greenhills.smart.com.ph:8080",
		},
		{
			URL:       "http://sp1.ss.zain.com/speedtest/upload.aspx",
			Latitude:  4.8594,
			Longitude: 31.5713,
			Name:      "Juba",
			Country:   "South Sudan",
			CC:        "SS",
			Sponsor:   "Zain South Sudan",
			ID:        "9011",
			URL2:      "http://sp2.ss.zain.com/speedtest/upload.aspx",
			Host:      "sp1.ss.zain.com:8080",
		},
		{
			URL:       "http://speedtest.directv.com.co/speedtest/upload.php",
			Latitude:  4.7110,
			Longitude: -74.0721,
			Name:      "Bogota",
			Country:   "Colombia",
			CC:        "CO",
			Sponsor:   "DIRECTV Colombia",
			ID:        "14393",
			Host:      "speedtest.directv.com.co:8080",
		},
		{
			URL:       "http://speedtest.creolink.com/speedtest/upload.php",
			Latitude:  4.0500,
			Longitude: 9.7000,
			Name:      "Douala",
			Country:   "Cameroon",
			CC:        "CM",
			Sponsor:   "CREOLINK COMMUNICATIONS",
			ID:        "9806",
			URL2:      "http://speedtest2.creolink.com/speedtest/upload.php",
			Host:      "speedtest.creolink.com:8080",
		},
		{
			URL:       "http://speedtest.camtel.cm/speedtest/upload.php",
			Latitude:  3.8480,
			Longitude: 11.5021,
			Name:      "Yaounde",
			Country:   "Cameroon",
			CC:        "CM",
			Sponsor:   "CAMTEL",
			ID:        "8634",
			URL2:      "http://speed.camtel.cm/speedtest/upload.php",
			Host:      "speedtest.camtel.cm:8080",
		},
		{
			URL:       "http://speedtest.ytlcomms.my/speedtest/upload.php",
			Latitude:  3.1357,
			Longitude: 101.6880,
			Name:      "Kuala Lumpur",
			Country:   "Malaysia",
			CC:        "MY",
			Sponsor:   "Yes 4G",
			ID:        "1701",
			URL2:      "http://183.78.1.15/speedtest/upload.php",
			Host:      "speedtest.ytlcomms.my:8080",
		},
		{
			URL:       "http://speedtest-ix.idola.net.id/speedtest/upload.php",
			Latitude:  1.3521,
			Longitude: 103.8198,
			Name:      "Singapore",
			Country:   "Singapore",
			CC:        "SG",
			Sponsor:   "PT. Aplikanusa Lintasarta",
			ID:        "14364",
			Host:      "speedtest-ix.idola.net.id:8080",
		},
		{
			URL:       "http://speedtest.ipi9.com/speedtest/upload.php",
			Latitude:  0.3944,
			Longitude: 9.4625,
			Name:      "Libreville",
			Country:   "Gabon",
			CC:        "GA",
			Sponsor:   "iPi9",
			ID:        "3793",
			URL2:      "http://speedtest.ipi9.com/speedtest2/upload.php",
			Host:      "speedtest.ipi9.com:8080",
		},
		{
			URL:       "http://speedtest.mtn.co.ug/upload.php",
			Latitude:  0.3167,
			Longitude: 32.5833,
			Name:      "Kampala",
			Country:   "Uganda",
			CC:        "UG",
			Sponsor:   "MTN Uganda",
			ID:        "3439",
			Host:      "speedtest.mtn.co.ug:8080",
		},
		{
			URL:       "http://speedtest.telkomkenya.co.ke/speedtest/upload.php",
			Latitude:  -1.2833,
			Longitude: 36.8167,
			Name:      "Nairobi",
			Country:   "Kenya",
			CC:        "KE",
			Sponsor:   "Telkom Kenya Ltd",
			ID:        "10914",
			URL2:      "http://speed.telkomkenya.co.ke/speedtest/upload.php",
			Host:      "speedtest.telkomkenya.co.ke:8080",
		},
		{
			URL:       "http://sp1.logicpro.com.br/speedtest/upload.php",
			Latitude:  -3.1000,
			Longitude: -60.0167,
			Name:      "Manaus",
			Country:   "Brazil",
			CC:        "BR",
			Sponsor:   "Logic Pro Tecnologia",
			ID:        "7449",
			URL2:      "http://sp2.logicpro.com.br/speedtest/upload.php",
			Host:      "sp1.logicpro.com.br:8080",
		},
		{
			URL:       "http://spetstmns01.timbrasil.com.br/speedtest/upload.php",
			Latitude:  -3.1000,
			Longitude: -60.0167,
			Name:      "Manaus",
			Country:   "Brazil",
			CC:        "BR",
			Sponsor:   "TIM Brasil",
			ID:        "12562",
			Host:      "spetstmns01.timbrasil.com.br:8080",
		},
		{
			URL:       "http://speedtest.cybernet.co.tz/speedtest/upload.php",
			Latitude:  -3.3667,
			Longitude: 36.6833,
			Name:      "Arusha",
			Country:   "Tanzania, United Republic of",
			CC:        "TZ",
			Sponsor:   "Arusha Art Ltd",
			ID:        "15888",
			Host:      "speedtest.cybernet.co.tz:8080",
		},
		{
			URL:       "http://sl-01-mba.ke.seacomnet.com/upload.php",
			Latitude:  -4.0500,
			Longitude: 39.6667,
			Name:      "Mombasa",
			Country:   "Kenya",
			CC:        "KE",
			Sponsor:   "Seacom Ltd",
			ID:        "5910",
			URL2:      "http://speedtest-mba.seacom.mu/upload.php",
			Host:      "sl-01-mba.ke.seacomnet.com:8080",
		},
		{
			URL:       "http://ookla-kin.vodanet.cd/speedtest/upload.php",
			Latitude:  -4.3250,
			Longitude: 15.3222,
			Name:      "Kinshasa",
			Country:   "DR Congo",
			CC:        "CD",
			Sponsor:   "Vodacom DRC",
			ID:        "6646",
			URL2:      "http://ookla1.vodanet.cd/speedtest/upload.php",
			Host:      "ookla-kin.vodanet.cd:8080",
		},
		{
			URL:       "http://ooklapnr.cg.airtel.com/speedtest/upload.php",
			Latitude:  -4.7692,
			Longitude: 11.8664,
			Name:      "Pointe Noire",
			Country:   "Congo",
			CC:        "CG",
			Sponsor:   "AIRTEL Congo B.",
			ID:        "15629",
			Host:      "ooklapnr.cg.airtel.com:8080",
		},
		{
			URL:       "http://jakarta.speedtest.telkom.net.id/speedtest/upload.php",
			Latitude:  -6.1745,
			Longitude: 106.8227,
			Name:      "Jakarta",
			Country:   "Indonesia",
			CC:        "ID",
			Sponsor:   "PT. Telekomunikasi Indonesia",
			ID:        "7582",
			URL2:      "http://jakarta.speedtest.telkom.co.id/speedtest/upload.php",
			Host:      "jakarta.speedtest.telkom.net.id:8080",
		},
		{
			URL:       "http://testspeed.tigo.co.tz/upload.php",
			Latitude:  -6.8167,
			Longitude: 39.2839,
			Name:      "Dar es Salaam",
			Country:   "Tanzania",
			CC:        "TZ",
			Sponsor:   "MIC Tanzania Ltd",
			ID:        "6110",
			URL2:      "http://speedtest.tigo.co.tz/upload.php",
			Host:      "speedtest.tigo.co.tz:8080",
		},
		{
			URL:       "http://speedtest1.angolatelecom.com/speedtest/upload.php",
			Latitude:  -8.8308,
			Longitude: 13.2450,
			Name:      "Luanda",
			Country:   "Angola",
			CC:        "AO",
			Sponsor:   "Angola Telecom",
			ID:        "4722",
			URL2:      "http://speedtest2.angolatelecom.com/speedtest/upload.php",
			Host:      "speedtest1.angolatelecom.com:8080",
		},
		{
			URL:       "http://drone.bmobile.com.pg/speedtest/upload.php",
			Latitude:  -9.5000,
			Longitude: 147.1167,
			Name:      "Port Moresby",
			Country:   "Papua New Guinea",
			CC:        "PG",
			Sponsor:   "Bmobile Vodafone Ltd",
			ID:        "9023",
			URL2:      "http://speed.bmobile.com.pg/speedtest/upload.php",
			Host:      "drone.bmobile.com.pg:8080",
		},
		{
			URL:       "http://speedtestfl.orange.mu/speedtest/upload.php",
			Latitude:  -20.1644,
			Longitude: 57.5041,
			Name:      "Floreal",
			Country:   "Mauritius",
			CC:        "MU",
			Sponsor:   "Mauritius Telecom Ltd",
			ID:        "3827",
			URL2:      "http://speedtestfl2.orange.mu/speedtest/upload.php",
			Host:      "speedtestfl.orange.mu:8080",
		},
		{
			URL:       "http://durban.spdtst.saix.net/speedtest/upload.php",
			Latitude:  -29.8697,
			Longitude: 31.0236,
			Name:      "Durban",
			Country:   "South Africa",
			CC:        "ZA",
			Sponsor:   "Telkom SA (SAIX)",
			ID:        "1881",
			Host:      "durban.spdtst.saix.net:8080",
		},
		{
			URL:       "http://speedtest-per1.node1.com.au/upload.php",
			Latitude:  -31.9505,
			Longitude: 115.8605,
			Name:      "Perth",
			Country:   "Australia",
			CC:        "AU",
			Sponsor:   "NODE1 Internet",
			ID:        "15097",
			Host:      "speedtest-per1.node1.com.au:8080",
		},
		{
			URL:       "http://speedtest.datalinksrl.com.ar/speedtest/upload.php",
			Latitude:  -32.5834,
			Longitude: -61.1673,
			Name:      "Totoras",
			Country:   "Argentina",
			CC:        "AR",
			Sponsor:   "Datalink S.R.L. - VAYNET",
			ID:        "14054",
			Host:      "speedtest.datalinksrl.com.ar:8080",
		},
		{
			URL:       "http://speedtest.tlink.cl/speedtest/upload.php",
			Latitude:  -33.4378,
			Longitude: -70.6504,
			Name:      "Santiago",
			Country:   "Chile",
			CC:        "CL",
			Sponsor:   "TLINK SpA",
			ID:        "14485",
			Host:      "speedtest.tlink.cl:8080",
		},
		{
			URL:       "http://speedtest.pfr.vodacombusiness.co.za/speedtest/upload.php",
			Latitude:  -33.7139,
			Longitude: 25.5207,
			Name:      "Port Elizabeth",
			Country:   "South Africa",
			CC:        "ZA",
			Sponsor:   "Vodacom (Pty) Ltd",
			ID:        "11493",
			Host:      "speedtest.pfr.vodacombusiness.co.za:8080",
		},
		{
			URL:       "http://Syd.optusnet.com.au/speedtest/upload.php",
			Latitude:  -33.8600,
			Longitude: 151.2111,
			Name:      "Sydney",
			Country:   "Australia",
			CC:        "AU",
			Sponsor:   "'Yes' Optus",
			ID:        "1267",
			URL2:      "http://s1.Syd.optusnet.com.au/speedtest/upload.php",
			Host:      "Syd.optusnet.com.au:8080",
		},
		{
			URL:       "http://stcpt.clearaccess.co.za/speedtest/upload.php",
			Latitude:  -33.9248,
			Longitude: 18.4240,
			Name:      "Cape Town",
			Country:   "South Africa",
			CC:        "ZA",
			Sponsor:   "Clear Access",
			ID:        "13922",
			Host:      "stcpt.clearaccess.co.za:8080",
		},
		{
			URL:       "http://velocidad.telecentro.net.ar/speedtest/upload.php",
			Latitude:  -34.6036,
			Longitude: -58.3817,
			Name:      "Buenos Aires",
			Country:   "Argentina",
			CC:        "AR",
			Sponsor:   "Telecentro",
			ID:        "900",
			URL2:      "http://190.55.63.89/speedtest/upload.php",
			Host:      "velocidad.telecentro.net.ar:8080",
		},
		{
			URL:       "http://speedtest1.convergia.com.ar/upload.php",
			Latitude:  -34.6036,
			Longitude: -58.3817,
			Name:      "Buenos Aires",
			Country:   "Argentina",
			CC:        "AR",
			Sponsor:   "Convergia Argentina",
			ID:        "11480",
			Host:      "speedtest1.convergia.com.ar:8080",
		},
		{
			URL:       "http://speedtest-auckland.telecom.co.nz/speedtest/upload.php",
			Latitude:  -36.8500,
			Longitude: 174.7833,
			Name:      "Auckland",
			Country:   "New Zealand",
			CC:        "NZ",
			Sponsor:   "Spark",
			ID:        "4135",
			Host:      "speedtest-auckland.telecom.co.nz:8080",
		},
		{
			URL:       "http://speedtest.dolomitesnetwork.it/speedtest/upload.php",
			Latitude:  46.6092,
			Longitude: 11.8942,
			Name:      "Badia",
			Country:   "Italy",
			CC:        "IT",
			Sponsor:   "Dolomites Network",
			ID:        "16096",
			Host:      "speedtest.dolomitesnetwork.it:8080",
		},
		{
			URL:       "http://sp1.swissnetwork.ch/speedtest/upload.php",
			Latitude:  46.5167,
			Longitude: 6.6333,
			Name:      "Lausanne",
			Country:   "Switzerland",
			CC:        "CH",
			Sponsor:   "Swiss Network SA",
			ID:        "7483",
			URL2:      "http://sp2.swissnetwork.ch/speedtest/upload.php",
			Host:      "sp1.swissnetwork.ch:8080",
		},
		{
			URL:       "http://paris1.speedtest.orange.fr/upload.php",
			Latitude:  48.8742,
			Longitude: 2.3470,
			Name:      "Paris",
			Country:   "France",
			CC:        "FR",
			Sponsor:   "Orange",
			ID:        "5559",
			URL2:      "http://paris1.speedtest-orange.hivane.net/speedtest/upload.php",
			Host:      "paris1.speedtest.orange.fr:8080",
		},
		{
			URL:       "http://31.172.233.6/speedtest/upload.php",
			Latitude:  43.2964,
			Longitude: 5.3700,
			Name:      "Marseille",
			Country:   "France",
			CC:        "FR",
			Sponsor:   "Orange",
			ID:        "4661",
			URL2:      "http://speedtest.13webhosting.com/speedtest/upload.php",
			Host:      "speedtest.13webhosting.com:8080",
		},
		{
			URL:       "http://speedtest.bcn.adamo.es/speedtest/upload.php",
			Latitude:  41.3857,
			Longitude: 2.1699,
			Name:      "Barcelona",
			Country:   "Spain",
			CC:        "ES",
			Sponsor:   "Adamo",
			ID:        "1695",
			Host:      "speedtest.bcn.adamo.es:8080",
		},
		{
			URL:       "http://speedtest-skg.greekstream.net/speedtest/upload.php",
			Latitude:  40.6500,
			Longitude: 22.9000,
			Name:      "Thessaloniki",
			Country:   "Greece",
			CC:        "GR",
			Sponsor:   "Greekstream Networks",
			ID:        "10980",
			Host:      "speedtest-skg.greekstream.net:8080",
		},
		{
			URL:       "http://a.lisboa.speedtest.net.zon.pt/speedtest/upload.php",
			Latitude:  38.7000,
			Longitude: -9.1833,
			Name:      "Lisbon",
			Country:   "Portugal",
			CC:        "PT",
			Sponsor:   "NOS",
			ID:        "1249",
			URL2:      "http://b.lisboa.speedtest.net.zon.pt/speedtest/upload.php",
			Host:      "a.lisboa.speedtest.net.zon.pt:8080",
		},
	}
}
