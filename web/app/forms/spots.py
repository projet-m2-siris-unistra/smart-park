from sanic_wtf import SanicForm
from wtforms import StringField, SubmitField, HiddenField, SelectField
from wtforms.validators import DataRequired, Length

from app.forms.widgets import BXInput, BXSelect, BXSubmit



class ConfigurationForm(SanicForm):

    coordinates = StringField(
        widget=BXInput(input_type="text"),
        label='Coordonnés',
        description='Les coordonnées de la place.',
        validators=[DataRequired(), Length(max=38)]
    )

    type = SelectField(
        widget=BXSelect(),
        label='Type',
        description='Changer le type de véhicule pour lequel sert cette place.',
        validators=[DataRequired()],
        choices=[
            ('car', 'Voiture'),
            ('bike', 'Moto'),
            ('truck', 'Poids Lourd')
        ]
    )

    device = SelectField(
        widget=BXSelect(),
        label='Capteur',
        description='Changer le capteur qui sera associé à cette place.',
        validators=[DataRequired()],
        coerce=int
    )

    submit = SubmitField(
        widget=BXSubmit(),
        label='Enregistrer'
    )
