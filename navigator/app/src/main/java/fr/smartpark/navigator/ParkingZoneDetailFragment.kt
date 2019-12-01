package fr.smartpark.navigator


import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.observe
import androidx.navigation.fragment.navArgs
import androidx.transition.TransitionInflater
import androidx.transition.TransitionManager
import fr.smartpark.navigator.databinding.FragmentParkingZoneDetailBinding
import fr.smartpark.navigator.utilities.InjectorUtils
import fr.smartpark.navigator.viewmodels.ParkingZoneDetailViewModel

class ParkingZoneDetailFragment : Fragment() {
    private lateinit var binding: FragmentParkingZoneDetailBinding

    private val args: ParkingZoneDetailFragmentArgs by navArgs()

    private val viewModel: ParkingZoneDetailViewModel by viewModels {
        InjectorUtils.provideParkingZoneDetailViewModelFactory(
            requireContext(),
            args.parkingZone.zoneId
        )
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        sharedElementEnterTransition =
            TransitionInflater.from(context).inflateTransition(R.transition.parking_detail)
        postponeEnterTransition()
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentParkingZoneDetailBinding.inflate(inflater, container, false)

        viewModel.zone.observe(viewLifecycleOwner) { zone ->
            binding.zone = zone
        }

        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding.apply {
            // Fallback to zone passed in args
            zone = viewModel.zone.value ?: args.parkingZone
            TransitionManager.beginDelayedTransition(zoneCard)
            executePendingBindings()
            startPostponedEnterTransition()
        }
    }
}
