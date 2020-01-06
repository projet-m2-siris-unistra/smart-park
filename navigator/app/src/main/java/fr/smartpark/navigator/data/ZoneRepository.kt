package fr.smartpark.navigator.data

import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.switchMap
import fr.smartpark.navigator.api.ZoneRemoteDataSource
import fr.smartpark.navigator.utilities.resultLiveData
import javax.inject.Inject

class ZoneRepository @Inject constructor(private val dao: ZoneDao,
                                         private val remoteDataSource: ZoneRemoteDataSource) {
    val tenantId = MutableLiveData<Long>()
    fun getZone(tenantId: Long, id: Long) = dao.getZone(tenantId, id)
    fun getZones(tenantId: Long) =
        resultLiveData(
            databaseQuery = { dao.getZones(tenantId) },
            networkCall = { remoteDataSource.getZones(tenantId) },
            saveCallResult = { dao.insertAll(it) }
        )
}