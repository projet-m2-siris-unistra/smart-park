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
        #res = Request.getTenant(12)
        self.id = tenant_id
        self.name = "Schmilbligheim"
        self.coordinates = {'coordinates': [7.7475, 48.5827]}
        self.zones = []

    '''
        NATS: topic = nom de la fonction
              payload = parametres en JSON
              communication en JSON 

        exemple réponse sur une place (inclu objet device et zone)
        {"place_id": 12}
        {
            "place_id": 12,
            "zone": {
                "zone_id": 4,
                ...
            },
            "device": {
                "device_id": 42,
                "status": "occupied",
                
            }
        } 
            ==> extension sur une couche (remonte pas jusqu'au tenant)

        # Send a request and expect a single response
        # and trigger timeout if not faster than 1 second.
        try:
            response = await nc.request("help", b'help me', timeout=1)
            print("Received response: {message}".format(
                message=response.data.decode()))
        except ErrTimeout:
            print("Request timed out")

    '''

    def getZones(self):
        # request the zones linked to this town
        zone = ZoneManagement("CENTRE")
        zoneList = [zone]
        return zoneList 

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
        self.desc = "Parking description"
        self.type = "Payant"
        self.color = "#f4e628"
        self.spots = []

    def toJson(self):
        return {
            "id" : self.id,
            "name" : self.name,
            "nb_total_spots" : self.nb_total_spots,
            "nb_taken_spots" : self.nb_taken_spots,
            "desc" : self.desc,
            "type" : self.type,
            "color" : self.color,
            "coordinates" : self.getPolygon(),
            "spots" : Tooling.jsonList(self.getSpotList())
        }

    # Getter / Setter #

    def getNbTotalSpots(self):
        return self.nb_total_spots

    def getNbTakenSpots(self):
        return self.nb_taken_spots
    

    # Data requests #

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

    def __init__(self):
        self.id = 12
        self.battery = 100 # between 0 and 100
        self.state = "free"
        # creation date
        # updated date
