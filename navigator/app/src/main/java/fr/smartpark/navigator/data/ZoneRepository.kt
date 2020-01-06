package fr.smartpark.navigator.data

import fr.smartpark.navigator.api.ZoneRemoteDataSource
import fr.smartpark.navigator.utilities.resultLiveData
import javax.inject.Inject

class ZoneRepository @Inject constructor(private val dao: ZoneDao,
                                         private val remoteDataSource: ZoneRemoteDataSource) {
    fun getZone(id: Long) = dao.getZone(id)
    val zones = resultLiveData(
        databaseQuery = { dao.getZones() },
        networkCall = { remoteDataSource.getZones() },
        saveCallResult = { dao.insertAll(it.zones) }
    )
}