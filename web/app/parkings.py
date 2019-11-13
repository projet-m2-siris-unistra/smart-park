import json

# Instance of a zone
class ZoneManagement:

    def __init__(self, nameArg):
        id = 1
        name = nameArg
        nb_total_spots = 456
        nb_taken_spots = 123
    
    # Data requests #

    def getPolygon(self):
        return json.dumps({
        'id': 'zone-polygon',
            'type': 'fill',
            'source': {
                'type': 'geojson',
                'data': { 
                    'type': 'Feature',
                    'geometry': {
                        'type': 'Polygon',
                        'coordinates': [[
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
                        ]]
                    }
                }
            },
            'paint': {
                'fill-color': '#f4e628',
                'fill-opacity': 0.2    
            }
        })
    
    def getSpotList(self):
        # requesting all spots belonging to this zone
        # SpotManagement spotList = []
        # loop for parsing all spots
        spot = SpotManagement()
        spotList[0] = spot1.getPoint()
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
        
        id = 112535
        name = "GARE1#124"
        available = True
        inService = True
        pointJson = this.getPoint()


    # Data requests # 

    def getPoint(self):
        return json.dumps({
            'type': 'Feature',
            'geometry': {
                'type': 'Point',
                'coordinates': [7.7475, 48.5827]
            },
            'properties': {
                'title': 'Parking#001',
                'description': 'État du parking: OK'
            }
        })


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

