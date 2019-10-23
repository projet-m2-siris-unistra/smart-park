
class ZoneManagement:
    def __init__(self):
        id = 1
        name = "GARE1"
        nb_total_spots = 456
        nb_taken_spots = 123

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



class SpotManagement:
    def __init__(self):
        id = 112535
        name = "GARE1#124"
        available = True
        inService = True

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

