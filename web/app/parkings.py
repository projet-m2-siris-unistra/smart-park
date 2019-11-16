from nats.aio.client import Client as NATS
#from app import app.nc as nc

# Instance of a tenant
class TenantManagement:

    def __init__(self, tenant_id):
        self.id = tenant_id
        #infos = self.getTenantFromDB()
        self.name = "Schmilbligheim"
        self.coordinates = {'coordinates': [7.7475, 48.5827]}

    async def msg_handler(msg):
        subject = msg.subject
        reply = msg.reply
        data = msg.data.decode()
        print("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))

    #def getTenantFromDB():
    #    sid = await nc.suscribe("postgre", cb=msg_handler)
    #    await nc.publish("postgre", b'tenant')


    def getZones(self):
        zone_list = []
        # request the zones linked to this town
        zone1 = ZoneManagement("CENTRE")
        zone_list.append(zone1)
        return zone_list

    def getTotalSpots(self):
        count = 0
        zoneList = self.getZones()
        for zone in zoneList:
            count += zone.getNbTotalSpots()
        return count
    
    def getTakenSpots(self):
        count = 0
        zoneList = self.getZones()
        for zone in zoneList:
            count += zone.getNbTakenSpots()
        return count


# Instance of a zone
class ZoneManagement:

    def __init__(self, nameArg):
        self.id = 1
        self.name = nameArg
        self.nb_total_spots = 456
        self.nb_taken_spots = 123

    # Getter / Setter #

    def getNbTotalSpots(self):
        return self.nb_total_spots

    def getNbTakenSpots(self):
        return self.nb_taken_spots
    

    # Data requests #

    def getPolygon(self):
        return {
            'coordinates': [
                [7.739396,48.579816],[7.742014,48.579957],
                [7.744117,48.579134],[7.747464,48.578623],
                [7.74888,48.57885],[7.751756,48.579929],
                [7.755189,48.581831],[7.756906,48.583251],
                [7.754288,48.58555],[7.753558,48.586061],
                [7.751455,48.586743],[7.748537,48.58714],
                [7.746906,48.586828],[7.744503,48.585834],
                [7.740769,48.584244],[7.73901,48.582967],
                [7.738409,48.581973],[7.738495,48.580781],
                [7.739396,48.579816]
            ],
            'color': '#f4e628',
        }
    
    def getSpotList(self):
        # requesting all spots belonging to this zone
        # loop for parsing all spots
        spot = SpotManagement()
        spotList = [spot]
        return spotList 


     # Statistics #

    def getDailyStats(self):
        stats = {
            'stats_type':'Quotidienne',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats
    
    def getWeeklyStats(self):
        stats = {
            'stats_type':'Hebdomadaire',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats

    def getMonthlyStats(self):
        stats = {
            'stats_type':'Mensuelle',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats

    def getAnnualStats(self):
        stats = {
            'stats_type':'Annuelle',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats
    
    def getAllStats(self):
        stats = []
        stats.append(self.getDailyStats())
        stats.append(self.getWeeklyStats())
        stats.append(self.getMonthlyStats())
        stats.append(self.getAnnualStats())

        return stats



# Instance of a parking spot 
class SpotManagement:

    def __init__(self):
        self.id = 124
        self.name = "CENTRE-124"
        self.state = "free"
        self.pointJson = self.getPoint()
        self.coordinates = {7.7475, 48.5827}

    def toJson(self):
        return {
            "id" : self.id,
            "name" : self.name,
            "state" : self.state,
            "point" : self.pointJson
        }

    # Data requests # 

    def getPoint(self):
        return {
            'type': 'Feature',
            'geometry': {
                'type': 'Point',
                'coordinates': [7.7475, 48.5827]
            }
        }


    # Statistics #

    def getDailyStats(self):
        stats = {
            'nbUsers':12
            }
        return stats
    
    def getWeeklyStats(self):
        stats = {
            'nbUsers':54
            }
        return stats

    def getMonthlyStats(self):
        stats = {
            'nbUsers':144
            }
        return stats

    def getAnnualStats(self):
        stats = {
            'nbUsers':1029
            }
        return stats

