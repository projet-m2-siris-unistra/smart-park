package fr.smartpark.navigator.data.models

import android.graphics.Color
import android.os.Parcelable
import androidx.room.ColumnInfo
import androidx.room.Entity
// import androidx.room.Ignore
import androidx.room.PrimaryKey
import com.google.gson.annotations.SerializedName
import kotlinx.android.parcel.Parcelize

@Parcelize
@Entity(tableName = "zones")
data class Zone(
    @PrimaryKey @ColumnInfo(name = "id") @SerializedName("zone_id") val zoneId: Long,
    @ColumnInfo(name = "tenant_id") @SerializedName("tenant_id") val tenantId: Long,
    val name: String,
    val type: String,
    val color: String?
    // @Ignore val geo: List<Point>?
) : Parcelable {
    fun parseColor(): Int? = color?.let {
        when (it.length) {
            6 -> Color.parseColor("#$it")
            7 -> Color.parseColor("#$it")
            else -> null
        }
    }
}