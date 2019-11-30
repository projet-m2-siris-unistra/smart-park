package fr.smartpark.navigator


import android.os.Bundle
import android.transition.TransitionInflater
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.navArgs
import fr.smartpark.navigator.databinding.FragmentParkingZoneDetailBinding

class ParkingZoneDetailFragment : Fragment() {
    private lateinit var binding: FragmentParkingZoneDetailBinding

    private val args: ParkingZoneDetailFragmentArgs by navArgs()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        sharedElementEnterTransition =
            TransitionInflater.from(context).inflateTransition(android.R.transition.move)
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // TODO: rebind to a LiveData<ParkingZone>
        binding = FragmentParkingZoneDetailBinding.inflate(inflater, container, false)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding.apply {
            zone = args.parkingZone
            executePendingBindings()
        }
    }
}
