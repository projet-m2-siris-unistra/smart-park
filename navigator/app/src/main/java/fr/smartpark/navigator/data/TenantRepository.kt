package fr.smartpark.navigator.data

import fr.smartpark.navigator.api.TenantRemoteDataSource
import fr.smartpark.navigator.utilities.resultLiveData
import javax.inject.Inject

class TenantRepository @Inject constructor(private val dao: TenantDao,
                                           private val remoteDataSource: TenantRemoteDataSource) {
    fun getTenant(tenantId: Long) = dao.getTenant(tenantId)
    val tenants = resultLiveData(
        databaseQuery = { dao.getTenants() },
        networkCall = { remoteDataSource.getTenants() },
        saveCallResult = { dao.insertAll(it) }
    )
}
