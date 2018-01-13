import re
from setuptools import setup
import ast

setup(
    name='Juicy',
    keywords=['database', 'client', 'distributed-system'],
    version='0.0.2',
    description='Juicy database python client',
    author='aljun',
    author_email='salamer_gaga@163.com',
    license='Apache',
    url='https://github.com/salamer/Juicy/python',
    download_url='https://github.com/salamer/Juicy/python',

    install_requires=[
        'grpcio',
    ],

    packages=['Juicy'],

    classifiers=[
        'Development Status :: 3 - Alpha',
        "License :: OSI Approved :: Apache Software License",
        'Environment :: Web Environment',
        "Programming Language :: Python :: 2.7",
        'Topic :: Internet :: WWW/HTTP',
        'Topic :: Software Development :: Libraries :: Python Modules'
    ]
)
