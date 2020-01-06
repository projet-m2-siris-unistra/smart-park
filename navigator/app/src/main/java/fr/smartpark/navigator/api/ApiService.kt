package fr.smartpark.navigator.api

import fr.smartpark.navigator.data.models.TenantListResponse
import fr.smartpark.navigator.data.models.ZoneListResponse
import retrofit2.Response
import retrofit2.http.GET
import retrofit2.http.Path
import retrofit2.http.Query

interface ApiService {
    @GET("tenants")
    suspend fun listTenants(
        @Query("limit") limit: Long = 100,
        @Query("offset") offset: Long? = null): Response<TenantListResponse>

    @GET("tenants/{tenant}/zones")
    suspend fun listZones(
        @Path("tenant") tenant: Long,
        @Query("limit") limit: Long = 100,
        @Query("offset") offset: Long? = null): Response<ZoneListResponse>
}