package fr.smartpark.navigator.data.models

import android.os.Parcelable
import kotlinx.android.parcel.Parcelize

@Parcelize
data class Point(
    val lat: Float,
    val long: Float
) : Parcelable