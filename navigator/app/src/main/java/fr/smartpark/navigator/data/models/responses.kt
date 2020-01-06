package fr.smartpark.navigator.data.models

import com.google.gson.annotations.SerializedName

data class PageInfo(
    val total: Long,
    @SerializedName("has_next") val hasNext: Boolean,
    @SerializedName("has_prev") val hasPrev: Boolean,
    val limit: Long,
    val offset: Long
)

data class ZoneListResponse(
    val page: PageInfo,
    val zones: List<Zone>
)
