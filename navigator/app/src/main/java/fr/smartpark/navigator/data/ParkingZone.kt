package fr.smartpark.navigator.data

import android.os.Parcelable
import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey
import kotlinx.android.parcel.Parcelize

@Parcelize
@Entity(tableName = "zones")
data class ParkingZone(
    @PrimaryKey @ColumnInfo(name = "id") val zoneId: String,
    val name: String,
    val color: String
) : Parcelable