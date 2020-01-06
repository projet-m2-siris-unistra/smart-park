package fr.smartpark.navigator.api

import fr.smartpark.navigator.data.models.TenantListResponse
import fr.smartpark.navigator.data.models.ZoneListResponse
import retrofit2.Response
import retrofit2.http.GET
import retrofit2.http.Query

interface ApiService {
    @GET("zones")
    suspend fun listZones(
        @Query("limit") limit: Long = 10,
        @Query("offset") offset: Long? = null): Response<ZoneListResponse>

    @GET("tenants")
    suspend fun listTenants(
        @Query("limit") limit: Long = 10,
        @Query("offset") offset: Long? = null): Response<TenantListResponse>
}