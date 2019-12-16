from sanic_wtf import SanicForm
from wtforms import StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit

from app.parkings import TenantManagement



class BaseForm(SanicForm):
    name = StringField(
        widget=BXInput(input_type="text"),
        label='Nom',
        description='Ce nom unique identifiera la zone.',
        validators=[DataRequired(), Length(max=40)]
    )

    type = SelectField(
        widget=BXSelect(),
        label='Type de la zone',
        description='Indiquer le type de parking que contiendra la zone.',
        validators=[DataRequired()],
        choices=[
            ('free', 'Gratuit'),
            ('paid', 'Payant'),
            ('blue', 'Zone bleue')
        ]
    )

    color = StringField(
        widget=BXInput(input_type="color"),
        render_kw={'class' : ''},
        label='Couleur',
        description='Choisissez une ,couleur associée à la zone. Cela facilitera sa reconnaissance sur les cartes.',
        validators=[DataRequired()]
    )



# Creation form
class CreationForm(BaseForm):
    polygon = HiddenField()

    submit = SubmitField(
        widget=BXSubmit(),
        label='Suivant'
    )


# General configuration form
class ConfigurationForm(CreationForm):
    submit = SubmitField(
        widget=BXSubmit(),
        label='Enregistrer'
    )

    delete = SubmitField(
        widget=BXSubmit(),
        render_kw={'color': 'danger'},
        label="Supprimer"
    )


# The form for creating a spot/place
class SpotsAddingForm(SanicForm):

    deviceSelect = SelectField(
        widget=BXSelect(),
        label='Capteur',
        description='Choisissez le capteur qui sera associé à cette place.',
        validators=[DataRequired()],
        coerce=int
    )

    typeSelect = SelectField(
        widget=BXSelect(),
        label='Type',
        description='Choisissez le type de véhicule pour lequel sert cette place.',
        validators=[DataRequired()],
        choices=[
            ('car', 'Voiture'),
            ('bike', 'Moto'),
            ('truck', 'Poids Lourd')
        ]
    )

    coordinatesInput = StringField(
        widget=BXInput(input_type="text"),
        label='Coordonnés',
        description='Les coordonnés de la place remplis manuellement ou avec la carte ci-dessus.',
        validators=[DataRequired(), Length(max=36)]
    )

    submit = SubmitField(
        widget=BXSubmit(),
        label='Enregistrer'
    )