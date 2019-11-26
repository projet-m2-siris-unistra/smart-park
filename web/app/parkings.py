import json as js
import datetime

from app.bus import Request

class Tooling:

    # convert elements of list into Json
    @staticmethod
    def jsonList(arg):
        liste = []
        for item in arg:
            liste.append(js.dumps(item.toJson()))
        return liste



# Instance of a tenant
class TenantManagement:

    def __init__(self, tenant_id):
        self.id = tenant_id
        self.zones = []
        # default data to see if DB requests are doing well
        self.name = "NOT UPDATED"
        self.coordinates = [7.9726, 49.0310] # Altenstadt (FR,67)


    async def init(self, tenant_id):
        response = await Request.getTenant(tenant_id)
        data = js.loads(response)
        self.name = data['name']
        self.coordinates = data['geo']


    # Get the list of all the zones from this tenant
    async def getZones(self):
        # request the zones linked to this town
        zoneList = []
        reponse = await Request.getZones(self.id)
        data = js.loads(reponse)

        for item in data:
            obj = ZoneManagement(item['zone_id'])
            obj.staticInit(
                item['name'],
                item['type'],
                '#' + item['color'],
                item['geo']
            )
            self.zones.append(obj)


    def getTotalSpots(self):
        count = 0
        if self.zones == []:
            print("WARNING: tenant.zones empty")
            return -1
        for zone in self.zones:
            count += zone.getNbTotalSpots()
        return count


    def getTakenSpots(self):
        count = 0
        if self.zones == []:
            print("WARNING: tenant.zones empty")
            return -1
        for zone in self.zones:
            count += zone.getNbTakenSpots()
        return count


# Instance of a zone
class ZoneManagement:

    def __init__(self, zone_id):
        self.id = zone_id
        # some default data to reveal further failed init
        self.desc = "Parking description"
        self.spots = []

    def staticInit(self, name, type, color, polygon):
        self.name = name
        self.type = type
        self.color = color
        self.polygon = polygon

    async def init(self, zone_id):
        response = await Request.getZone(zone_id)
        data = js.loads(response)
        self.name = data['name']
        self.type = data['type']
        self.color = '#' + data['color']
        print("polygon from data: ", data['geo'])
        self.polygon = data['geo']

    def toJson(self):
        return {
            "id" : self.id,
            "name" : self.name,
            "nb_total_spots" : self.getNbTotalSpots(),
            "nb_taken_spots" : self.getNbTakenSpots(),
            "desc" : self.desc,
            "type" : self.type,
            "color" : self.color,
            "coordinates" : self.getPolygon(),
            "spots" : Tooling.jsonList(self.getSpotList())
        }

    # Getter / Setter #

    def getNbTotalSpots(self):
        # calculation from DB
        return 321

    def getNbTakenSpots(self):
        # calculation from DB
        return 123
    

    """
    def getPolygon(self):
        return [
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
        ]
    """
    
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
        self.coordinates = [7.7475, 48.5827]
        self.device = "lul"

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
            'stats_type':'Daily',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats
    
    def getWeeklyStats(self):
        stats = {
            'stats_type':'Weekly',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats

    def getMonthlyStats(self):
        stats = {
            'stats_type':'Monthly',
            'total_users':123,
            'rate':18,
            'is_charge':True,
            'avg_price':0.34,
            'earning':12
            }
        return stats

    def getAnnualStats(self):
        stats = {
            'stats_type':'Annual',
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



# instance of a device
class DeviceManagement:

    def __init__(self, device_id):
        self.id = device_id
        self.battery = -1 # between 0 and 100
        self.state = "free"
        # creation date
        # updated date

    async def init(self, device_id):
        response = await Request.getDevice(device_id)
        data = js.loads(response)
        self.battery